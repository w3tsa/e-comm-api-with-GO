package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID 			int16	`json:"id"`
	Name 		string	`json:"name"`
	ImageUrl 	string `json:"imageUrl"`
	Price 		float64	`json:"price"`
	Category 	string `json:"category"`
}


var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	products = append(products, Product{ID: 1, Name: "Skinny Jeans", ImageUrl: "url", Price: 44.36, Category: "Pants"})
	products = append(products, Product{ID: 2, Name: "Regular Jeans", ImageUrl: "url", Price: 44.36, Category: "Pants"})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}