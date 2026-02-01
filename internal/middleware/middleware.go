package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs HTTP method, path, and timestamp for each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request
		log.Printf("Request: %s %s at %s", r.Method, r.URL.Path, start.Format(time.RFC3339))

		// Call the next handler
		next.ServeHTTP(w, r)

	})
}

func DurationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next.ServeHTTP(w, r)

		// Log the duration
		duration := time.Since(start)
		log.Printf("Completed: %s %s in %v", r.Method, r.URL.Path, duration)
	})
}
