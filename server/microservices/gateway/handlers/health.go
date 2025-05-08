package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

// Health represents the health status of the server
type Health struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
	Memory    struct {
		Alloc      uint64 `json:"alloc"`
		TotalAlloc uint64 `json:"totalAlloc"`
		Sys        uint64 `json:"sys"`
		NumGC      uint32 `json:"numGC"`
	} `json:"memory"`
	Version string `json:"version"`
}

var startTime = time.Now()

// HealthCheckHandler returns information about the server's health
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	health := Health{
		Status:    "UP",
		Timestamp: time.Now(),
		Uptime:    time.Since(startTime).String(),
		Version:   runtime.Version(),
	}

	// Add memory stats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	health.Memory.Alloc = memStats.Alloc / 1024 / 1024      // MB
	health.Memory.TotalAlloc = memStats.TotalAlloc / 1024 / 1024 // MB
	health.Memory.Sys = memStats.Sys / 1024 / 1024          // MB
	health.Memory.NumGC = memStats.NumGC

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(health)
}

// ReadinessCheckHandler checks if the server is ready to accept requests
func ReadinessCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "READY"})
}

// LivenessCheckHandler checks if the server is alive
func LivenessCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ALIVE"})
}

