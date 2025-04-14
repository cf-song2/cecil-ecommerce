package service

import (
	"cecil-ecommerce/internal/model"
	"cecil-ecommerce/internal/repository"
)

func GetAllProducts() ([]model.Product, error) {
	return repository.GetAllProducts()
}

func GetProductByID(id int) (*model.Product, error) {
	return repository.GetProductByID(id)
}
