package containersController

import (
	"net/http"
	"encoding/json"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
)

func ListContainers(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, "Failed to create Docker client", http.StatusInternalServerError)
		return
	}

	// Get the list of running containers
	containers, err := cli.ContainerList(r.Context(), types.ContainerListOptions{})
	if err != nil {
		http.Error(w, "Failed to retrieve container list", http.StatusInternalServerError)
		return
	}

	// Create a slice to store container information
	var containerInfo []types.Container

	// Iterate over the containers and extract relevant information
	for _, container := range containers {
		containerInfo = append(containerInfo, types.Container{
			ID:    container.ID,
			Names: container.Names,
			Image: container.Image,
			State: container.State,
		})
	}

	// Convert the container information to JSON
	jsonData, err := json.Marshal(containerInfo)
	if err != nil {
		http.Error(w, "Failed to marshal container info to JSON", http.StatusInternalServerError)
		return
	}

	// Set the response content type
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}
