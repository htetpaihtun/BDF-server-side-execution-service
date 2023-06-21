package users

import (
	"net/http"
)

// Handler for the /users route
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users page"))
}
