package http

import (
	"net/http"

	"github.com/febriaricandra/go-habit-tracker/internal/http/handlers"
	"github.com/febriaricandra/go-habit-tracker/internal/http/middleware"
)

// register user, habit, and activity routes
func AuthRoutes(mux *http.ServeMux) {

	mux.Handle("/api/user", middleware.ChainMiddleware(http.HandlerFunc(
		handlers.GetUserByName),
		middleware.MethodAllow("GET"),
		middleware.MiddlewareAuthJWT))
}

func PublicRoutes(mux *http.ServeMux) {
	mux.Handle("/api/user/register", middleware.ChainMiddleware(http.HandlerFunc(
		handlers.RegisterHandler),
		middleware.MethodAllow("POST")))

	mux.Handle("/api/user/login", middleware.ChainMiddleware(http.HandlerFunc(
		handlers.LoginHandler),
		middleware.MethodAllow("POST")))

	mux.Handle("/api/", middleware.ChainMiddleware(
		http.HandlerFunc(handlers.HomeHandler),
		middleware.MethodAllow("GET")))
}
