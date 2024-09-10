package handlers

import (
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User registered successfully"))
}
