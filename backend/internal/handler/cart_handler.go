package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cecil-ecommerce/internal/service"
	"cecil-ecommerce/internal/util"
	"cecil-ecommerce/internal/repository"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) Handle(w http.ResponseWriter, r *http.Request) {
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
		h.handleGetCart(w, user.ID)
	case http.MethodPost:
		h.handleAddItem(w, r, user.ID)
	case http.MethodPut:
		h.handleUpdateItem(w, r, user.ID)
	case http.MethodDelete:
		h.handleDeleteItem(w, r, user.ID)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CartHandler) handleGetCart(w http.ResponseWriter, userID int) {
	cart, err := h.service.GetCart(userID)
	if err != nil {
		http.Error(w, "Failed to fetch cart", http.StatusInternalServerError)
		return
	}
	util.JSON(w, http.StatusOK, cart)
}

func (h *CartHandler) handleAddItem(w http.ResponseWriter, r *http.Request, userID int) {
	var item struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.service.AddItem(userID, item.ProductID, item.Quantity); err != nil {
		http.Error(w, "Failed to add to cart", http.StatusInternalServerError)
		return
	}
	util.JSON(w, http.StatusOK, "Item added to cart")
}

func (h *CartHandler) handleUpdateItem(w http.ResponseWriter, r *http.Request, userID int) {
	var item struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateItem(userID, item.ProductID, item.Quantity); err != nil {
		http.Error(w, "Failed to update cart item", http.StatusInternalServerError)
		return
	}
	util.JSON(w, http.StatusOK, "Cart item updated")
}

func (h *CartHandler) handleDeleteItem(w http.ResponseWriter, r *http.Request, userID int) {
	productIDStr := r.URL.Query().Get("product_id")
	if productIDStr != "" {
		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}
		if err := h.service.RemoveItem(userID, productID); err != nil {
			http.Error(w, "Failed to remove item", http.StatusInternalServerError)
			return
		}
		util.JSON(w, http.StatusOK, "Item removed")
		return
	}

	if err := h.service.ClearCart(userID); err != nil {
		http.Error(w, "Failed to clear cart", http.StatusInternalServerError)
		return
	}
	util.JSON(w, http.StatusOK, "Cart cleared")
}
