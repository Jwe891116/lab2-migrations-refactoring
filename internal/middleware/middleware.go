// Package declaration - defines middleware package
package middleware

//Import necessary packages
import (
	"log"      //For logging request details and durations
	"net/http" //For http handler interface
	"time"     //For timing requests and formatting timestamps
)

// LoggingMiddleware logs HTTP method, path, and timestamp for each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Records the time the request was made
		//Used for log purposes
		start := time.Now()

		// Log the request
		log.Printf("Request: %s %s at %s", r.Method, r.URL.Path, start.Format(time.RFC3339))

		// Call the next handler
		next.ServeHTTP(w, r)

	})
}

// DurationMiddleware measures and logs request processing time
func DurationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Records the start time of the request
		start := time.Now()

		//Call the next handler
		next.ServeHTTP(w, r)

		//Calculate the duration since request started
		duration := time.Since(start)
		//Log the completion with duration
		log.Printf("Completed: %s %s in %v", r.Method, r.URL.Path, duration)
	})
}
