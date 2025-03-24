package database

import (
    "database/sql"
    "log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/jackc/pgx/v5/stdlib" // Import pgx driver for database/sql compatibility
)

var DB *sql.DB

func ConnectDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading environment variables")
    }

    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL is not set in environment variables")
    }

    var errOpen error
    DB, errOpen = sql.Open("pgx", dsn) // Use pgx with database/sql
    if errOpen != nil {
        log.Fatalf("Error connecting to Postgres: %v", errOpen)
    }

    if err := DB.Ping(); err != nil {
        log.Fatalf("Database not reachable: %v", err)
    }

    log.Println("Connected to PostgreSQL successfully!")
}

func CloseDB() {
    if DB != nil {
        err := DB.Close()
        if err != nil {
            log.Println("Error closing database:", err)
        } else {
            log.Println("Database connection closed.")
        }
    }
}
