package main

import (
	"log"
	"net/http"

	"cecil-ecommerce/internal/config"
	"cecil-ecommerce/internal/handler"
	"cecil-ecommerce/internal/repository"
	"cecil-ecommerce/internal/service"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		allowedOrigins := map[string]bool{
			"https://cecil-personal.site":     true,
			"https://www.cecil-personal.site": true,
			"https://api.cecil-personal.site": true,
		}

		if allowedOrigins[origin] {
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

	// Repository Layer
	repos := &repository.Repositories{
		Product: repository.NewProductRepository(db),
		User:    repository.NewUserRepository(db),
		Cart:    repository.NewCartRepository(db),
	}

	// Service Layer
	authService := service.NewAuthService(repos.User)
	productService := service.NewProductService(repos.Product)
	cartService := service.NewCartService(repos.Cart)

	// Handler Layer
	authHandler := handler.NewAuthHandler(authService)
	productHandler := handler.NewProductHandler(productService)
	cartHandler := handler.NewCartHandler(cartService)

	// Route Binding
	http.HandleFunc("/api/login", withCORS(authHandler.Login))
	http.HandleFunc("/api/register", withCORS(authHandler.Register))
	http.HandleFunc("/api/me", withCORS(authHandler.Me))
	http.HandleFunc("/api/products", withCORS(productHandler.List))
	http.HandleFunc("/api/product", withCORS(productHandler.Detail))
	http.HandleFunc("/api/cart", withCORS(cartHandler.Handle))

	log.Printf("[server] listening on %s", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
