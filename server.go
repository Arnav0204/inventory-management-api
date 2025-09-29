package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Server struct {
	DB *sql.DB
}

func (srv *Server) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var productRequest Product
	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&productRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert into DB, RETURNING product_id
	err = srv.DB.QueryRow(
		`INSERT INTO products (product_name, product_quantity, product_description)
         VALUES ($1, $2, $3) RETURNING product_id`,
		productRequest.ProductName, productRequest.ProductQuantity, productRequest.ProductDescription,
	).Scan(&productRequest.ProductId)
	if err != nil {
		http.Error(w, "DB insert failed", http.StatusInternalServerError)
		return
	}

	// Return the product with the generated ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productRequest)
}

func (srv *Server) GetProductHandler(w http.ResponseWriter, r *http.Request) {

}

func (srv *Server) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {

}

func (srv *Server) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

}
