package imagesHandler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/images/list":
		ListImages(w, r)
	// case "/containers/create":
	// 	CreateContainer(w, r)
	default:
		http.NotFound(w, r)
	}
}
