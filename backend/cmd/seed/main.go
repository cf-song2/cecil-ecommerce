package main

import (
	"cecil-ecommerce/internal/config"
	"cecil-ecommerce/internal/repository"
	"cecil-ecommerce/internal/service"
	"log"
)

func main() {
	cfg := config.Load()
	db := config.ConnectDB(cfg.DBUrl)
	defer db.Close()

	repo := repository.NewProductRepository(db)
	seeder := service.NewProductSeeder(repo)

	if err := seeder.SeedFromAPI("https://fakestoreapi.com/products"); err != nil {
		log.Fatalf("[Failed] Failed to seed products: %v", err)
	}
}
