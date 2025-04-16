package handler

import (
	"net/http"
	"strconv"

	"cecil-ecommerce/internal/service"
	"cecil-ecommerce/internal/util"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to load products", http.StatusInternalServerError)
		return
	}
	util.JSON(w, http.StatusOK, products)
}

func (h *ProductHandler) Detail(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	util.JSON(w, http.StatusOK, product)
}
