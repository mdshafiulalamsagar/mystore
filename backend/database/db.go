package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

// DB is the global database instance
var DB *sql.DB

// ConnectDB initializes the PostgreSQL connection using environment variables
func ConnectDB() {
    // 1. Render/Neon-er cloud DATABASE_URL check korbe
    dbURL := os.Getenv("DATABASE_URL")

    // 2. Cloud URL na thakle local environment variables use korbe
    if dbURL == "" {
        host := os.Getenv("DB_HOST")
        port := os.Getenv("DB_PORT")
        user := os.Getenv("DB_USER")
        password := os.Getenv("DB_PASSWORD")
        dbname := os.Getenv("DB_NAME")

        if host == "" { host = "localhost" }
        if port == "" { port = "5432" }
        if user == "" { user = "postgres" }
        if dbname == "" { dbname = "mystore_db" }

        dbURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
            host, port, user, password, dbname)
    }

    var err error
    DB, err = sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatalf("Failed to open database connection: %v", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    fmt.Println("Successfully connected to the PostgreSQL database!")
}