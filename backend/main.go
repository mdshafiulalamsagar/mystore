package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/mdshafiulalamsagar/mystore/backend/config"
    "github.com/mdshafiulalamsagar/mystore/backend/database"
    "github.com/mdshafiulalamsagar/mystore/backend/routes"
)

// enableCORS middleware allows cross-origin requests from React frontend
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight OPTIONS request
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func main() {
    config.LoadConfig()
    database.ConnectDB()

    // Register all routes
    routes.SetupRoutes()

    // Dynamic Port detection for Render
    port := config.Port
    if port == "" {
        port = os.Getenv("PORT")
    }
    if port == "" {
        port = "8080"
    }

    fmt.Printf("Server is running on port %s...\n", port)

    // Wrap default http mux with CORS middleware
    log.Fatal(http.ListenAndServe(":"+port, enableCORS(http.DefaultServeMux)))
}