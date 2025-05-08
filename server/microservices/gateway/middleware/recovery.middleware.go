package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"go.uber.org/zap"
)

// RecoveryMiddleware captures panics and recovers to prevent the server from crashing
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the stack trace
				stackTrace := debug.Stack()
				logger.GetLogger().Error("Server panic recovered",
					zap.Any("error", err),
					zap.String("stack", string(stackTrace)),
					zap.String("path", r.URL.Path),
					zap.String("method", r.Method),
					zap.String("remote_addr", r.RemoteAddr),
				)

				// Return an error to the client
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

