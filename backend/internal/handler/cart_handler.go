package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cecil-ecommerce/internal/repository"
	"cecil-ecommerce/internal/service"
	"cecil-ecommerce/internal/util"
)

func CartHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := util.GetSessionID(r)
	if err != nil || sessionID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case http.MethodGet:
		cart, err := service.GetCart(user.ID)
		if err != nil {
			http.Error(w, "Failed to fetch cart", http.StatusInternalServerError)
			return
		}
		util.JSON(w, http.StatusOK, cart)

	case http.MethodPost:
		var item struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		}
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		if err := service.AddItemToCart(user.ID, item.ProductID, item.Quantity); err != nil {
			http.Error(w, "Failed to add to cart", http.StatusInternalServerError)
			return
		}
		util.JSON(w, http.StatusOK, "Item added to cart")

	case http.MethodPut:
		var item struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		}
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		if err := service.UpdateCartItem(user.ID, item.ProductID, item.Quantity); err != nil {
			http.Error(w, "Failed to update cart item", http.StatusInternalServerError)
			return
		}
		util.JSON(w, http.StatusOK, "Cart item updated")

	case http.MethodDelete:
		// support ?product_id=X to remove specific item
		productIDStr := r.URL.Query().Get("product_id")
		if productIDStr != "" {
			productID, err := strconv.Atoi(productIDStr)
			if err != nil {
				http.Error(w, "Invalid product ID", http.StatusBadRequest)
				return
			}
			if err := service.RemoveItemFromCart(user.ID, productID); err != nil {
				http.Error(w, "Failed to remove item", http.StatusInternalServerError)
				return
			}
			util.JSON(w, http.StatusOK, "Item removed")
			return
		}

		// otherwise clear entire cart
		if err := service.ClearUserCart(user.ID); err != nil {
			http.Error(w, "Failed to clear cart", http.StatusInternalServerError)
			return
		}
		util.JSON(w, http.StatusOK, "Cart cleared")

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
