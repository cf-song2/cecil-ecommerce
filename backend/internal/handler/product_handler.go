package handler

import (
	"net/http"
	"strconv"

	"cecil-ecommerce/internal/service"
	"cecil-ecommerce/internal/util"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := service.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to load products", http.StatusInternalServerError)
		return
	}
	util.JSON(w, http.StatusOK, products)
}

func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := service.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	util.JSON(w, http.StatusOK, product)
}
