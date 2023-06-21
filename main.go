package main

import (
	"log"
	"net/http"
	"github.com/htetpaihtun/BDF-server-side-execution-service/containerController"
)

func main() {
	// Create a new router
	router := http.NewServeMux()

	// Define the routes and their corresponding handlers
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/container", container-controller.Handler)

	// Start the HTTP server
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler for the home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}
