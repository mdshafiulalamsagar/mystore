package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/routes"
)

// main is the entry point of the application
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize database connection
	database.ConnectDB()

	// Register API routes
	routes.SetupRoutes()

	// Define default HTTP route for health check
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to myStore Backend API!")
	})

	// Get port from environment variables or set default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is starting on port :%s\n", port)

	// Start server and handle failure
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}