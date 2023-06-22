package logger

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"io/ioutil"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func WriteLog(filePath string) error {
	
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	// Subscribe to Docker events
	ctx := context.Background()
	eventOptions := types.EventsOptions{}
	events, errors := cli.Events(ctx, eventOptions)

	// Handle errors in a separate goroutine
	go func() {
		for err := range errors {
			log.Println("Docker event error:", err)
		}
	}()

	// Process Docker events and write to file in real-time
	for event := range events {
		eventJSON, err := json.Marshal(event)
		if err != nil {
			log.Println("Failed to marshal event:", err)
			continue
		}
		_, err = file.Write(append(eventJSON, []byte("\n")...))
		if err != nil {
			log.Println("Failed to write event to file:", err)
			continue
		}
		file.Sync()
	}

	return nil
}

func RetrieveLog(w http.ResponseWriter, r *http.Request) {

	logFilePath := "./logger/logs/docker-logs.log" // log dir should be dynamic : TO FIX LATER

	// Read the log file
	logData, err := ioutil.ReadFile(logFilePath)
	if err != nil {
		http.Error(w, "Failed to read log file", http.StatusInternalServerError)
		return
	}

	// Set the response content type
	w.Header().Set("Content-Type", "text/plain")

	// Write the log data to the response
	w.Write(logData)
}