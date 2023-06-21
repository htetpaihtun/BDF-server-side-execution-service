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
	router.HandleFunc("/container", containerController.Handler)
	router.HandleFunc("/containers", containerController.CheckContainers)

	// Start the HTTP server
	log.Println("Server listening on :9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}

// Handler for the home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Service is up and running.. C:"))
}
