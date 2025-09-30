package main

type Product struct {
	ProductId          int64  `json:"product_id"`
	ProductName        string `json:"product_name"`
	ProductQuantity    int64  `json:"product_quantity"`
	ProductDescription string `json:"product_description"`
}
