package imagesHandler

import (
	"net/http"
	"encoding/json"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
)

func ListImages(w http.ResponseWriter, r *http.Request) {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, "Failed to create Docker client", http.StatusInternalServerError)
		return
	}

	// Get the list of running containers
	images, err := cli.ImageList(r.Context(), types.ImageListOptions{})
	if err != nil {
		http.Error(w, "Failed to retrieve image list", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(images, "", "    ")
	if err != nil {
		http.Error(w, "Failed to marshal image info to JSON", http.StatusInternalServerError)
		return
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the HTTP response writer
	w.Write(jsonData)
}

