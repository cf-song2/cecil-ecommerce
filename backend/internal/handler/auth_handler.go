package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"cecil-ecommerce/internal/repository"
	"cecil-ecommerce/internal/service"
	"cecil-ecommerce/internal/util"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	creds.Username = strings.TrimSpace(creds.Username)
	creds.Password = strings.TrimSpace(creds.Password)

	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Username and password required", http.StatusBadRequest)
		return
	}

	log.Printf("[LoginHandler] Attempting login for: '%s'", creds.Username)

	user, sessionID, err := service.AuthenticateUser(creds.Username, creds.Password)
	if err != nil {
		log.Printf("[LoginHandler] ❌ Auth failed for '%s': %v", creds.Username, err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	util.SetSessionCookie(w, sessionID)

	log.Printf("[LoginHandler] ✅ Auth success for '%s', session: %s", creds.Username, sessionID)

	json.NewEncoder(w).Encode(map[string]any{
		"message": "logged in",
		"user":    user.Username,
	})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if creds.Username == "" || creds.Password == "" || creds.Email == "" {
		http.Error(w, "All fields required", http.StatusBadRequest)
		return
	}

	if err := repository.CreateUser(creds.Username, creds.Email, creds.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.JSON(w, http.StatusCreated, "User registered")
}

func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := util.GetSessionID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	util.JSON(w, http.StatusOK, user)
}
