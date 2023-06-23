package containersHandler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/containers/list":
		ListContainers(w, r)
	case "/containers/create":
		CreateContainer(w, r)
	default:
		http.NotFound(w, r)
	}
}
