package routes

import (
	"net/http"

	"github.com/Jwe891116/lab2-JoshuaEmmanuel/internal/handlers"
	"github.com/Jwe891116/lab2-JoshuaEmmanuel/internal/middleware"
)

// SetupRoutes configures all routes with middleware and returns the mux
func SetupRoutes(mux *http.ServeMux) *http.ServeMux {
	// Apply global middleware (logging & durations for all routes)
	mux.Handle("/", middleware.LoggingMiddleware(middleware.DurationMiddleware(http.HandlerFunc(handlers.Home))))
	mux.Handle("/about", middleware.LoggingMiddleware(middleware.DurationMiddleware(http.HandlerFunc(handlers.Home))))
	mux.Handle("/contact", middleware.LoggingMiddleware(middleware.DurationMiddleware(http.HandlerFunc(handlers.Home))))
	mux.Handle("/calculate", middleware.LoggingMiddleware(middleware.DurationMiddleware(http.HandlerFunc(handlers.Home))))

	return mux
}
