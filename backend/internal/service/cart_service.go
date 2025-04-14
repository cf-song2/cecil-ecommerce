package service

import (
	"cecil-ecommerce/internal/model"
	"cecil-ecommerce/internal/repository"
)

func AddItemToCart(userID, productID, quantity int) error {
	return repository.AddToCart(userID, productID, quantity)
}

func GetCart(userID int) ([]model.CartItem, error) {
	return repository.GetCartItems(userID)
}

func UpdateCartItem(userID, productID, quantity int) error {
	return repository.UpdateCartItem(userID, productID, quantity)
}

func RemoveItemFromCart(userID, productID int) error {
	return repository.RemoveCartItem(userID, productID)
}

func ClearUserCart(userID int) error {
	return repository.ClearCart(userID)
}
