package http

import (
	"net/http"

	"github.com/febriaricandra/go-habit-tracker/internal/http/handlers"
	"github.com/febriaricandra/go-habit-tracker/internal/http/middleware"
)

// register user, habit, and activity routes
func AuthRoutes(mux *http.ServeMux) {
	// mux.HandleFunc("POST /api/user/register", handlers.CreateUserHandler)
	// mux.HandleFunc("GET /api/", handlers.HomeHandler)

	mux.Handle("/api/user/register", middleware.ChainMiddleware(http.HandlerFunc(
		handlers.CreateUserHandler),
		middleware.MethodAllow("POST"),
		middleware.MiddlewareAuthJWT))

	mux.Handle("/api/user/login", middleware.ChainMiddleware(http.HandlerFunc(
		handlers.LoginHandler),
		middleware.MethodAllow("POST")))

	mux.Handle("/api/users", middleware.ChainMiddleware(http.HandlerFunc(
		handlers.GetUsersHandler),
		middleware.MethodAllow("GET"),
		middleware.MiddlewareAuthJWT))

	mux.Handle("/api/", middleware.ChainMiddleware(
		http.HandlerFunc(handlers.HomeHandler),
		middleware.MethodAllow("GET")))
}
