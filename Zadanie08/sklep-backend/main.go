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
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		products := []Product{
			{ID: 1, Name: "Jablko", Price: 1.2},
			{ID: 2, Name: "Banan", Price: 1.5},
		}
		if err := json.NewEncoder(w).Encode(products); err != nil {
                http.Error(w, "Blad JSONa", http.StatusInternalServerError)
            }
	})

	http.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == "POST" {
			var data map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
                    http.Error(w, "Bad request", http.StatusBadRequest)
                    return
                }
			fmt.Printf("Otrzymano platnosc.)")
			w.WriteHeader(http.StatusCreated)
		}
	})

	fmt.Println("URL backendu: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
            fmt.Printf("Blad uruchamiania serwera: %v\n", err)
        }
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}