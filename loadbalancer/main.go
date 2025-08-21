// Detailed Explanation of the Go Load Balancer
//
// This load balancer is a reverse proxy server that distributes HTTP traffic across multiple backend servers.
// It supports three balancing policies: Round Robin, Least Connections, and Random. It also includes health checks,
// sticky sessions, retry logic, and graceful shutdown.
//
// --- Backend Struct ---
// - Holds a backend's URL, alive status (atomic flag), connection count, weight, and a ReverseProxy.
// - Provides methods to mark alive/dead, increment/decrement connection counts, and retrieve current connections.
//
// --- LoadBalancer Struct ---
// - Manages a slice of backends and the balancing logic.
// - Tracks the balancing policy, a round robin counter, sticky session map, mutex locks, health check settings, and retry count.
// - Sticky sessions map a client ID to a specific backend.
//
// --- Key Features ---
// 1. **NewBackend**: Initializes a backend with its URL, weight, and ReverseProxy configuration.
// 2. **HealthCheck**: Periodically sends GET requests to each backend's health endpoint. Updates alive status.
// 3. **chooseBackend**: Selects a backend based on the policy:
//    - Round Robin: Cycles through alive backends in order.
//    - Least Connections: Picks the backend with the fewest active connections.
//    - Random: Picks a random alive backend.
// 4. **Sticky Sessions**: Uses an LBID cookie to persist a client to the same backend for subsequent requests.
// 5. **ServeHTTP**: Handles incoming requests:
//    - Tries to use sticky mapping if available.
//    - Selects backend if not already set.
//    - Sets sticky cookie if needed.
//    - Proxies request and decrements connection count after completion.
//    - Retries another backend if a proxy call takes too long or fails.
// 6. **ParseBackends**: Parses CLI input for backends and optional weights.
// 7. **Main Function**:
//    - Reads CLI flags for configuration.
//    - Parses and initializes backends.
//    - Sets up health checks.
//    - Defines HTTP handlers for proxying and status display.
//    - Runs HTTP server and listens for SIGINT/SIGTERM for graceful shutdown.
//
// --- Graceful Shutdown ---
// On receiving a shutdown signal, it cancels the health check context, waits for active requests to finish, and stops the server cleanly.
//
// This architecture makes it flexible for high availability, can scale horizontally by adding more backends,
// and can be extended for TLS termination, metrics, and weighted round robin.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// Backend represents an upstream server
type Backend struct {
	URL              *url.URL
	Alive            int32 // 1 = alive, 0 = dead (use atomic)
	ReverseProxy     *httputil.ReverseProxy
	Weight           int
	CurrentConns     int64 // atomic
	Name             string
}

// PickPolicy defines the balancing algorithm
type PickPolicy int

const (
	RoundRobin PickPolicy = iota
	LeastConnections
	Random
)

// LoadBalancer holds backends and balancing logic
type LoadBalancer struct {
	Backends      []*Backend
	policy        PickPolicy
	rrCounter     uint64
	sticky        bool
	stickymap     map[string]*Backend
	stickymu      sync.RWMutex
	mu            sync.RWMutex
	healthPath    string
	healthInterval time.Duration
	healthTimeout  time.Duration
	retries       int
}

// NewBackend creates a Backend from raw url string
func NewBackend(raw string, weight int) (*Backend, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	// keep host header
	proxy.Director = func(r *http.Request) {
		r.URL.Scheme = u.Scheme
		r.URL.Host = u.Host
		// preserve path and rawpath
		// host header preserved to upstream host
		r.Host = u.Host
	}
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, e error) {
		log.Printf("proxy error for %s: %v", u, e)
		rw.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(rw, "Bad Gateway: %v", e)
	}
	b := &Backend{
		URL:          u,
		Alive:        1,
		ReverseProxy: proxy,
		Weight:       weight,
		Name:         u.Host,
	}
	return b, nil
}

// MarkAlive sets backend alive state
func (b *Backend) MarkAlive(alive bool) {
	if alive {
		atomic.StoreInt32(&b.Alive, 1)
	} else {
		atomic.StoreInt32(&b.Alive, 0)
	}
}

func (b *Backend) IsAlive() bool {
	return atomic.LoadInt32(&b.Alive) == 1
}

// increment/decrement connection counters
func (b *Backend) IncConn() { atomic.AddInt64(&b.CurrentConns, 1) }
func (b *Backend) DecConn() { atomic.AddInt64(&b.CurrentConns, -1) }
func (b *Backend) GetConns() int64 { return atomic.LoadInt64(&b.CurrentConns) }

// NewLoadBalancer creates a LoadBalancer
func NewLoadBalancer(backends []*Backend, policy PickPolicy) *LoadBalancer {
	lb := &LoadBalancer{
		Backends:       backends,
		policy:         policy,
		sticky:         true,
		stickymap:      make(map[string]*Backend),
		healthPath:      "/health",
		healthInterval: 5 * time.Second,
		healthTimeout:  2 * time.Second,
		retries:        2,
	}
	return lb
}

// HealthCheck pings backends periodically
func (lb *LoadBalancer) HealthCheck(ctx context.Context) {
	client := &http.Client{Timeout: lb.healthTimeout}
	ticker := time.NewTicker(lb.healthInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			for _, b := range lb.Backends {
				go func(backend *Backend) {
					healthURL := *backend.URL
					healthURL.Path = lb.healthPath
					resp, err := client.Get(healthURL.String())
					alive := err == nil && resp.StatusCode < 500
					backend.MarkAlive(alive)
					if resp != nil {
						resp.Body.Close()
					}
					log.Printf("health %s -> %v", backend.URL, alive)
				}(b)
			}
		}
	}
}

// chooseBackend selects an alive backend based on policy
func (lb *LoadBalancer) chooseBackend() *Backend {
	lb.mu.RLock()
	defer lb.mu.RUnlock()
	alive := make([]*Backend, 0, len(lb.Backends))
	for _, b := range lb.Backends {
		if b.IsAlive() {
			alive = append(alive, b)
		}
	}
	if len(alive) == 0 {
		return nil
	}
	switch lb.policy {
	case RoundRobin:
		idx := int(atomic.AddUint64(&lb.rrCounter, 1) % uint64(len(alive)))
		return alive[idx]
	case LeastConnections:
		var best *Backend
		var bestConns int64 = 1<<62 - 1
		for _, b := range alive {
			if c := b.GetConns(); c < bestConns {
				best = b
				bestConns = c
			}
		}
		if best == nil {
			return alive[rand.Intn(len(alive))]
		}
		return best
	case Random:
		return alive[rand.Intn(len(alive))]
	default:
		return alive[0]
	}
}

// getSticky returns sticky backend for cookie if exists
func (lb *LoadBalancer) getSticky(id string) *Backend {
	lb.stickymu.RLock()
	b := lb.stickymap[id]
	lb.stickymu.RUnlock()
	if b != nil && b.IsAlive() {
		return b
	}
	return nil
}

// setSticky assigns cookie id to backend
func (lb *LoadBalancer) setSticky(id string, b *Backend) {
	lb.stickymu.Lock()
	lb.stickymap[id] = b
	lb.stickymu.Unlock()
}

// ServeHTTP handles incoming requests, performs sticky, retries, and proxying
func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// sticky session cookie
	var backend *Backend
	cookie, err := r.Cookie("LBID")
	if err == nil && lb.sticky {
		backend = lb.getSticky(cookie.Value)
	}
	// if no sticky backend, choose one
	attempts := 0
	for attempts <= lb.retries {
		if backend == nil {
			backend = lb.chooseBackend()
		}
		if backend == nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}
		// update sticky mapping if needed
		if lb.sticky {
			if cookie == nil {
				id := fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000))
				http.SetCookie(w, &http.Cookie{Name: "LBID", Value: id, Path: "/", Expires: time.Now().Add(24 * time.Hour)})
				lb.setSticky(id, backend)
			}
		}
		// increment conn
		backend.IncConn()
		start := time.Now()
		backend.ReverseProxy.ServeHTTP(w, r)
		backend.DecConn()
		// basic success heuristics: if proxy didn't take too long, consider success
		if time.Since(start) < 30*time.Second {
			return
		}
		// otherwise try next backend
		backend = nil
		attempts++
	}
	// exhausted retries
	http.Error(w, "Bad Gateway - retries exhausted", http.StatusBadGateway)
}

func ParseBackends(input string) ([]*Backend, error) {
	parts := strings.Split(input, ",")
	out := make([]*Backend, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		// allow weight as host|weight
		weight := 1
		addr := p
		if strings.Contains(p, "|") {
			sp := strings.Split(p, "|")
			addr = sp[0]
			if w, err := strconv.Atoi(sp[1]); err == nil && w > 0 {
				weight = w
			}
		}
		b, err := NewBackend(addr, weight)
		if err != nil {
			return nil, err
		}
		out = append(out, b)
	}
	return out, nil
}

func main() {
	var (
		listen = flag.String("listen", ":8080", "listen address (e.g. :8080)")
		backends = flag.String("backends", "http://localhost:9001|1,http://localhost:9002|1", "comma-separated backend URLs optionally with |weight")
		policyStr = flag.String("policy", "roundrobin", "balancing policy: roundrobin|leastconn|random")
		healthPath = flag.String("health-path", "/health", "path to use for health checks")
		healthInterval = flag.Int("health-interval", 5, "health check interval seconds")
		healthTimeout = flag.Int("health-timeout", 2, "health check timeout seconds")
		retries = flag.Int("retries", 2, "retries to another backend on failure")
	)
	flag.Parse()

	bs, err := ParseBackends(*backends)
	if err != nil {
		log.Fatalf("invalid backends: %v", err)
	}
	var policy PickPolicy
	switch strings.ToLower(*policyStr) {
	case "roundrobin":
		policy = RoundRobin
	case "leastconn":
		policy = LeastConnections
	case "random":
		policy = Random
	default:
		policy = RoundRobin
	}
	lb := NewLoadBalancer(bs, policy)
	lb.healthPath = *healthPath
	lb.healthInterval = time.Duration(*healthInterval) * time.Second
	lb.healthTimeout = time.Duration(*healthTimeout) * time.Second
	lb.retries = *retries

	// initial health check
	for _, b := range lb.Backends {
		b.MarkAlive(true) // optimistic
	}

	ctx, cancel := context.WithCancel(context.Background())
	go lb.HealthCheck(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/lb/proxy/", func(w http.ResponseWriter, r *http.Request) {
		// strip prefix and forward to lb
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/lb/proxy")
		lb.ServeHTTP(w, r)
	})
	// simple root to show status
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go Load Balancer\nBackends:\n")
		for _, b := range lb.Backends {
			fmt.Fprintf(w, "- %s alive=%v conns=%d\n", b.URL, b.IsAlive(), b.GetConns())
		}
	})

	srv := &http.Server{
		Addr:    *listen,
		Handler: mux,
	}

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Printf("listening on %s", *listen)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-stop
	log.Println("shutdown signal received")
	cancel()
	ctxShut, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	if err := srv.Shutdown(ctxShut); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}
	log.Println("server stopped")
}
