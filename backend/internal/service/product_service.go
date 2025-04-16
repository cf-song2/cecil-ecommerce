package service

import (
	"cecil-ecommerce/internal/model"
	"cecil-ecommerce/internal/repository"
)

type ProductService interface {
	GetAll() ([]model.Product, error)
	GetByID(id int) (*model.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{repo: r}
}

func (s *productService) GetAll() ([]model.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetByID(id int) (*model.Product, error) {
	return s.repo.GetByID(id)
}
