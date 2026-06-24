package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	"log"
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/joho/godotenv"
)

var db *sql.DB

var googleOauthConfig = &oauth2.Config{
	RedirectURL: "http://localhost:8080/auth/google/callback",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint: google.Endpoint,
}

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

func generateStateOauthCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(20 * time.Minute)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)
	return state
}

func main() {
	initDB()
	defer db.Close()

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Błąd ładowania .env")
    }

	googleOauthConfig.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
    googleOauthConfig.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

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

		token := base64.StdEncoding.EncodeToString([]byte("local:" + creds.Username))

		response := map[string]string{
			"message": "Zalogowano pomyslnie",
			"token": token,
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


	http.HandleFunc("/auth/google/login", func(w http.ResponseWriter, r *http.Request) {
		oauthState := generateStateOauthCookie(w)
		u := googleOauthConfig.AuthCodeURL(oauthState)
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/auth/google/callback", func(w http.ResponseWriter, r *http.Request) {
		oauthState, _ := r.Cookie("oauthstate")
		if r.FormValue("state") != oauthState.Value {
			log.Println("Niezgodny stan OAuth (potencjalny atak CSRF)")
			http.Redirect(w, r, "http://localhost:5173/login?error=invalid_state", http.StatusTemporaryRedirect)
			return
		}

		token, err := googleOauthConfig.Exchange(context.Background(), r.FormValue("code"))
		if err != nil {
			log.Printf("Błąd wymiany tokenu: %v\n", err)
			http.Redirect(w, r, "http://localhost:5173/login?error=token_exchange_failed", http.StatusTemporaryRedirect)
			return
		}

		response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			log.Printf("Błąd pobierania danych użytkownika: %v\n", err)
			http.Redirect(w, r, "http://localhost:5173/login?error=user_info_failed", http.StatusTemporaryRedirect)
			return
		}
		defer response.Body.Close()

		contents, err := io.ReadAll(response.Body)
		if err != nil {
			log.Printf("Błąd odczytu danych: %v\n", err)
			http.Redirect(w, r, "http://localhost:5173/login?error=read_failed", http.StatusTemporaryRedirect)
			return
		}

		var user map[string]interface{}
		json.Unmarshal(contents, &user)

		email := user["email"].(string)

		appToken := base64.StdEncoding.EncodeToString([]byte("google:" + email))

		http.Redirect(w, r, "http://localhost:5173/login?token="+appToken, http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/user/me", func(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

    rawToken := r.Header.Get("Authorization")
    if rawToken == "" {
        http.Error(w, "Niezalogowany", http.StatusUnauthorized)
        return
    }

    decoded, err := base64.StdEncoding.DecodeString(rawToken)
    if err != nil {
        http.Error(w, "Błędny token", http.StatusUnauthorized)
        return
    }
    
	tokenContent := string(decoded)

    username := tokenContent
    if strings.HasPrefix(tokenContent, "google:") {
        username = strings.TrimPrefix(tokenContent, "google:")
    } else if strings.HasPrefix(tokenContent, "local:") {
        username = strings.TrimPrefix(tokenContent, "local:")
    }
    response := map[string]string{
        "username": username,
    }
    w.Header().Set("Content-Type", "application/json")
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
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}