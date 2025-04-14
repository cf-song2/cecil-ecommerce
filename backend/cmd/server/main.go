package main

import (
	"log"
	"net/http"

	"cecil-ecommerce/internal/config"
	"cecil-ecommerce/internal/handler"
	"cecil-ecommerce/internal/repository"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "http://localhost" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h(w, r)
	}
}

func main() {
	cfg := config.Load()
	db := config.ConnectDB(cfg.DBUrl)
	defer db.Close()

	repository.SetDB(db)

	http.HandleFunc("/api/login", withCORS(handler.LoginHandler))
	http.HandleFunc("/api/products", withCORS(handler.ProductsHandler))
	http.HandleFunc("/api/product", withCORS(handler.ProductDetailHandler))
	http.HandleFunc("/api/cart", withCORS(handler.CartHandler))
	http.HandleFunc("/api/register", withCORS(handler.RegisterHandler))
	http.HandleFunc("/api/me", withCORS(handler.CurrentUserHandler))

	log.Printf("[server] listening on %s", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
