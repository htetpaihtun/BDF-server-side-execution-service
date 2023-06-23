package imagesHandler

import (
	"net/http"
)


// Handler 
func BuildImage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Build image coming soon.."))
}
