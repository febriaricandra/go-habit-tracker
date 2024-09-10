package main

import (
	"log"
	"net/http"
	"os"

	"github.com/febriaricandra/go-habit-tracker/internal/database"
	route "github.com/febriaricandra/go-habit-tracker/internal/http"
	"github.com/joho/godotenv"
)

func main() {

	//default port
	defaultPort := "8080"
	godotenv.Load(".env")
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set || you need to setup .env file")
	}

	//initialize database connection
	database.Connect()
	database.Migrate()

	//routing
	mux := http.NewServeMux()
	route.RegisterUserRoutes(mux)

	//start server
	srv := &http.Server{
		Addr:    ":" + defaultPort,
		Handler: mux,
	}

	log.Printf("Server started on port %s", defaultPort)
	log.Fatal(srv.ListenAndServe())
}
