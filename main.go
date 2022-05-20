package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
	const PORT = ":8080"
	
	// Routes Handlers
	router :=  r.PathPrefix("/api").Subrouter()
    
	router.HandleFunc("/products", getProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(PORT, r))
}