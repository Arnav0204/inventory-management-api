package main

import (
	"database/sql"
	"encoding/json"
	"log"
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
	rows, err := srv.DB.Query(`SELECT product_id, product_name, product_quantity, product_description 
								 FROM products`)
	if err != nil {
		log.Fatal("unable to fetch rows from DB")
		http.Error(w, "DB query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		scanErr := rows.Scan(&product.ProductId, &product.ProductName, &product.ProductQuantity, &product.ProductDescription)
		if scanErr != nil {
			http.Error(w, "Row scan failed", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (srv *Server) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var productRequest Product

	err := json.NewDecoder(r.Body).Decode(&productRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if productRequest.ProductId == 0 {
		http.Error(w, "Missing product_id", http.StatusBadRequest)
		return
	}

	result, err := srv.DB.Exec(
		`UPDATE products 
         SET product_name = $1, product_quantity = $2, product_description = $3
         WHERE product_id = $4`,
		productRequest.ProductName, productRequest.ProductQuantity, productRequest.ProductDescription, productRequest.ProductId,
	)
	if err != nil {
		http.Error(w, "DB update failed", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to check update result", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Return updated product
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productRequest)

}

func (srv *Server) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

}
