package middleware

import (
	"net/http"
	"sync"

	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// RateLimiter is a struct that holds rate limiters for each client.
type RateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
	rate     rate.Limit
	burst    int
}

// NewRateLimiter creates a new RateLimiter instance.
func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
		rate:     r,
		burst:    b,
	}
}

// GetLimiter retrieves or creates a rate limiter for a specific client.
func (rl *RateLimiter) GetLimiter(clientIP string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.limiters[clientIP]
	if !exists {
		limiter = rate.NewLimiter(rl.rate, rl.burst)
		rl.limiters[clientIP] = limiter
	}

	return limiter
}

// RateLimitMiddleware is an HTTP middleware that applies rate limiting.
func RateLimitMiddleware(rl *RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr // You can use a more robust method to extract the client IP.
		endpoint := r.URL.Path
		limiter := rl.GetLimiter(clientIP + endpoint)
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			logger.GetLogger().Warn("Rate limit exceeded",
				zap.String("client_ip", clientIP),
				zap.String("endpoint", endpoint),
				zap.Int("status_code", http.StatusTooManyRequests))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ValidateRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && r.Method != http.MethodGet {
			http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
			return
		}

		// Add additional validation logic here
		next.ServeHTTP(w, r)
	})
}

func LimitRequestBodyMiddleware(maxBytes int64, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://0.0.0.0:7556")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
