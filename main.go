package main

import (
	"log"
	"bufio"
	"strings"
	"os"
	"net/http"
	"github.com/htetpaihtun/BDF-server-side-execution-service/containersController"
)

func main() {

	// Load env variable 
	err := loadEnvFromFile(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new router
	router := http.NewServeMux()

	// Define the routes and their corresponding handlers
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/containers/", containersController.Handler)

	// Start the HTTP server
	log.Println("Server listening on :9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}

// Handler for the home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Service is up and running.. C:"))
}

func loadEnvFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("Invalid line in %s: %s", filename, line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set the environment variable in the current process
		err := os.Setenv(key, value)
		if err != nil {
			log.Printf("Failed to set environment variable: %s", key)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}