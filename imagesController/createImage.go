package imagesController
// package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type DockerfileData struct {
	Image    string
	WorkDir  string
	// Files    []string
	Commands string
}

func main() {
	// Define the data for the Dockerfile
	data := DockerfileData{
		Image:    "alpine:latest",
		WorkDir:  "/app",
		// Files:    []string{"main.go", "config.yaml"},
		Commands: "mkdir 'here'",
	}

	// Read the template file
	tmplFile := "../DockerfileTemplates/Alpine.tmpl"
	templateContent, err := ioutil.ReadFile(tmplFile)
	if err != nil {
		panic(err)
	}

	// Create a new template and parse the template content
	t, err := template.New("Dockerfile").Parse(string(templateContent))
	if err != nil {
		panic(err)
	}

	// Open a new file for writing the Dockerfile
	file, err := os.Create("Dockerfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Execute the template with the data and write the output to the file
	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
	
	println("Dockerfile generated successfully!")

}