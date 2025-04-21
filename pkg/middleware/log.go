package middleware

import (
	"To_Do/pkg/logger"
	"net/http"
	"time"
)

func SlogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Logger.Info("Входящий запрос",
			"Method", r.Method,
			"Path", r.URL.Path,
			"Time", time.Since(start),
		)
	})
}
