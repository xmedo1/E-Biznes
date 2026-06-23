package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/payments", paymentsHandler)
	http.HandleFunc("/cart", cartHandler)

	fmt.Println("URL backendu: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	products := []Product{
		{ID: 1, Name: "Jablko", Price: 1.2},
		{ID: 2, Name: "Banan", Price: 1.5},
	}
	json.NewEncoder(w).Encode(products)
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == "POST" {
		var data map[string]interface{}
		json.NewDecoder(r.Body).Decode(&data)
		fmt.Printf("Otrzymano platnosc: %v\n", data)
		w.WriteHeader(http.StatusCreated)
	}
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == "POST" {
		var data []interface{}
		json.NewDecoder(r.Body).Decode(&data)
		fmt.Printf("Otrzymano koszyk: %v\n", data)
		w.WriteHeader(http.StatusCreated)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}