// Package declaration - defines as the routes package
// Contains http route functions that process web requests
package routes

//Import necessary packages
import (
	"net/http" //For ServeMux, Handler, and routing functionality

	// Internal package imports (following Go module path conventions)
	// These are local packages within the same module

	//Contains handler functions
	"github.com/Jwe891116/lab2-JoshuaEmmanuel/internal/handlers"
	//Contains middlerware functions
	"github.com/Jwe891116/lab2-JoshuaEmmanuel/internal/middleware"
)

// SetupRoutes configures all routes with middleware and returns the mux
func SetupRoutes(mux *http.ServeMux) *http.ServeMux {
	// Apply global middleware (logging & durations for all routes)
	mux.Handle("/", middleware.LoggingMiddleware(
		middleware.DurationMiddleware(
			http.HandlerFunc(handlers.Home))))

	mux.Handle("/about", middleware.LoggingMiddleware(
		middleware.DurationMiddleware(
			http.HandlerFunc(handlers.Home))))

	mux.Handle("/contact", middleware.LoggingMiddleware(
		middleware.DurationMiddleware(
			http.HandlerFunc(handlers.Home))))

	mux.Handle("/calculate", middleware.LoggingMiddleware(
		middleware.DurationMiddleware(
			http.HandlerFunc(handlers.Home))))

	return mux
}
