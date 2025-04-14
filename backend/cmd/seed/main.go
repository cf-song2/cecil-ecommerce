package main

import (
	"cecil-ecommerce/internal/config"
	"cecil-ecommerce/internal/model"
	"cecil-ecommerce/internal/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ProductAPI struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

func main() {
	cfg := config.Load()
	db := config.ConnectDB(cfg.DBUrl)
	repository.SetDB(db)
	defer db.Close()

	res, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		log.Fatalf("Failed to fetch products: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var apiProducts []ProductAPI
	if err := json.Unmarshal(body, &apiProducts); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	for _, p := range apiProducts {
		product := model.Product{
			Name:        p.Title,
			Description: p.Description,
			Price:       int(p.Price * 100),
			ImageURL:    p.Image,
		}

		if err := repository.CreateProductIfNotExists(product); err != nil {
			log.Printf("Failed to insert %s: %v", p.Title, err)
		} else {
			log.Printf("Inserted: %s", p.Title)
		}
	}
}
