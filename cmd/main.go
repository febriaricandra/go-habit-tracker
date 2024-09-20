package main

import (
	"log"
	"net/http"

	"github.com/febriaricandra/go-habit-tracker/config"
	"github.com/febriaricandra/go-habit-tracker/internal/database"
	route "github.com/febriaricandra/go-habit-tracker/internal/http"
	"github.com/rs/cors"
)

func main() {

	// load config
	cfg := config.LoadConfig()

	//default port
	defaultPort := "8080"

	//initialize database connection
	database.Connect(cfg)
	database.Migrate()

	//routing
	mux := http.NewServeMux()
	route.AuthRoutes(mux)
	route.PublicRoutes(mux)

	//CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(mux)

	//start server
	srv := &http.Server{
		Addr:    ":" + defaultPort,
		Handler: corsHandler,
	}

	log.Printf("Server started on port %s", defaultPort)
	log.Fatal(srv.ListenAndServe())
}
