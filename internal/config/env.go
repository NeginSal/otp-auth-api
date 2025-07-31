package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file once at startup
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or failed to load")
	}
}

// GetEnv gets value with fallback
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
