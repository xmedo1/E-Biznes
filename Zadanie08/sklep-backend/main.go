package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT UNIQUE NOT NULL,
		"password" TEXT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db.Exec(`INSERT OR IGNORE INTO users (username, password) VALUES ('admin', 'password')`)
}

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var testUser = map[string]string{
	"admin": "password",
}

func main() {
	initDB()
	defer db.Close()

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
			fmt.Printf("Otrzymano platnosc.")
			w.WriteHeader(http.StatusCreated)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		var creds Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Niepoprawny format danych", http.StatusBadRequest)
			return
		}

		var dbPassword string
		err := db.QueryRow("SELECT password FROM users WHERE username = ?", creds.Username).Scan(&dbPassword)
		
		if err == sql.ErrNoRows || dbPassword != creds.Password {
			http.Error(w, "Zle dane logowania", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "Blad serwera", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"message": "Zalogowano pomyslnie",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})


	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		var creds Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Niepoprawny format danych", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", creds.Username, creds.Password)
		
		if err != nil {
			http.Error(w, "Login zajety", http.StatusConflict)
			return
		}

		response := map[string]string{
			"message": "Zarejestrowano pomyslnie",
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
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