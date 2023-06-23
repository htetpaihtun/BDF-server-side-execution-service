package containersHandler

import (
	"net/http"
)

func CreateContainer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Smth"))
}
