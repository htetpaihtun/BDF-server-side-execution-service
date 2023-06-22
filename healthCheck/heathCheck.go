package healthCheck

import (
	"net/http"
	"fmt"
	
	// "github.com/docker/docker/client"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Perform health check logic here
	// You can check dependencies, database connections, etc.
	
	// cli, err := client.NewClientWithOpts(client.FromEnv)
	// if err != nil {
	// 	http.Error(w, "Failed to create Docker client", http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Service is healthy")
}