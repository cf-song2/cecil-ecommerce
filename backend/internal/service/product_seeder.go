package service

import (
	"cecil-ecommerce/internal/model"
	"cecil-ecommerce/internal/repository"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ProductAPI struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

type ProductSeeder struct {
	repo repository.ProductRepository
}

func NewProductSeeder(r repository.ProductRepository) *ProductSeeder {
	return &ProductSeeder{repo: r}
}

func (ps *ProductSeeder) SeedFromAPI(apiURL string) error {
	res, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("failed to fetch: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}

	var items []ProductAPI
	if err := json.Unmarshal(body, &items); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	for _, p := range items {
		product := model.Product{
			Name:        p.Title,
			Description: p.Description,
			Price:       int(p.Price * 100),
			ImageURL:    p.Image,
		}
		if err := ps.repo.CreateIfNotExists(product); err != nil {
			fmt.Printf("[Fail] Failed to insert %s: %v\n", p.Title, err)
		} else {
			fmt.Printf("[Success] Inserted: %s\n", p.Title)
		}
	}
	return nil
}
