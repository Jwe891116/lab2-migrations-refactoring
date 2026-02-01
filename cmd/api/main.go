package main

import (
	"log"
	"net/http"

	"github.com/Jwe891116/lab2-JoshuaEmmanuel/internal/routes"
)

func main() {
	// Create a new multiplexer
	mux := http.NewServeMux()

	// Setup routes with middleware
	mux = routes.SetupRoutes(mux)

	// Start the server
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
