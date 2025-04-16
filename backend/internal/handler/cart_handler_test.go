package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cecil-ecommerce/internal/model"
)

// It will fail
func TestCartHandler_POST(t *testing.T) {
	payload := model.CartItem{ProductID: 1, Quantity: 2}
	body, _ := json.Marshal(payload)
	r := httptest.NewRequest(http.MethodPost, "/api/cart", bytes.NewReader(body))
	w := httptest.NewRecorder()

	CartHandler(w, r)
	if w.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", w.Code)
	}
}

func TestCartHandler_GET(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/api/cart", nil)
	w := httptest.NewRecorder()

	CartHandler(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", w.Code)
	}
}
