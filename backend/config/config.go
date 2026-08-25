package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

// LoadConfig loads environment variables from .env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: No .env file found, reading system environment variables")
	}

	Port = getEnv("PORT", "8080")
	DBHost = getEnv("DB_HOST", "localhost")
	DBPort = getEnv("DB_PORT", "5432")
	DBUser = getEnv("DB_USER", "postgres")
	DBPassword = getEnv("DB_PASSWORD", "postgres")
	DBName = getEnv("DB_NAME", "mystore_db")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}