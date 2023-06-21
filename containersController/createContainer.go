package containersController

import (
	"net/http"
)


// Handler 
func CreateContainer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Smth"))
}
