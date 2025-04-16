package service

import (
	"cecil-ecommerce/internal/model"
	"cecil-ecommerce/internal/repository"
)

type CartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) AddItem(userID, productID, quantity int) error {
	return s.repo.Add(userID, productID, quantity)
}

func (s *CartService) GetCart(userID int) ([]model.CartItem, error) {
	return s.repo.Get(userID)
}

func (s *CartService) UpdateItem(userID, productID, quantity int) error {
	return s.repo.Update(userID, productID, quantity)
}

func (s *CartService) RemoveItem(userID, productID int) error {
	return s.repo.Remove(userID, productID)
}

func (s *CartService) ClearCart(userID int) error {
	return s.repo.Clear(userID)
}
