package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {

}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	var _ = initDB()
	r := mux.NewRouter()
	r.HandleFunc("/create-product", CreateProductHandler).Methods("POST")
	r.HandleFunc("/get-product", GetProductHandler).Methods("GET")
	r.HandleFunc("/update-product", UpdateProductHandler).Methods("POST")
	r.HandleFunc("/delete-product", DeleteProductHandler).Methods("DELETE")
	log.Println("all handlers registered")

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))

}
