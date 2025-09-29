package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load environment variables")
	}
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		log.Fatal("DB_URL is not set in environment")
	}

	connection, connectionErr := sql.Open("postgres", dbURL)
	if connectionErr != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	pingErr := connection.Ping()
	if pingErr != nil {
		log.Fatal("Cannot ping DB:", err)
	}

	log.Println("Database connected successfully")
	return connection
}

func main() {
	var _ = initDB()
	log.Println("starting go server")
}
