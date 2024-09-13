package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	JwtSecret string
	DB_USER   string
	DB_PASS   string
	DB_HOST   string
	DB_PORT   string
	DB_NAME   string
}

// LoadConfig loads the application configuration from the environment variables.
func LoadConfig() *AppConfig {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &AppConfig{
		JwtSecret: os.Getenv("JWT_SECRET"),
		DB_USER:   os.Getenv("DB_USER"),
		DB_PASS:   os.Getenv("DB_PASS"),
		DB_HOST:   os.Getenv("DB_HOST"),
		DB_PORT:   os.Getenv("DB_PORT"),
		DB_NAME:   os.Getenv("DB_NAME"),
	}
}

func GetSecret() string {
	return LoadConfig().JwtSecret
}
