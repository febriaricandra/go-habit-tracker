package http

import (
	"net/http"

	"github.com/febriaricandra/go-habit-tracker/internal/http/handlers"
)

// register user, habit, and activity routes
func RegisterUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user/register", handlers.CreateUserHandler)
}
