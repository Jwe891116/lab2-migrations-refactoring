// Package declaration - main execuable package
package main

//Import necessary packages for the application
import (
	"log"      //For logging messages and errors
	"net/http" //For http server fuctionality and routing

	//Internal package import - contains route definitions
	"github.com/Jwe891116/lab2-JoshuaEmmanuel/internal/routes"
)

func main() {
	// Create a new multiplexer
	//Matches incoming request URLs and calls appropriate handler function
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
