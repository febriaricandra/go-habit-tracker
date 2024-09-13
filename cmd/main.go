package main

import (
	"log"
	"net/http"

	"github.com/febriaricandra/go-habit-tracker/config"
	"github.com/febriaricandra/go-habit-tracker/internal/database"
	route "github.com/febriaricandra/go-habit-tracker/internal/http"
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

	//start server
	srv := &http.Server{
		Addr:    ":" + defaultPort,
		Handler: mux,
	}

	log.Printf("Server started on port %s", defaultPort)
	log.Fatal(srv.ListenAndServe())
}
