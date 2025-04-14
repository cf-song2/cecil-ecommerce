package service

import (
	"cecil-ecommerce/internal/model"
	"cecil-ecommerce/internal/repository"
	"cecil-ecommerce/internal/util"
	"errors"
	"log"
	"strings"
)

func AuthenticateUser(username, password string) (model.User, string, error) {
	userPtr, err := repository.GetUserByUsername(username)
	if err != nil {
		log.Printf("[AuthService] No user found: %s", username)
		return model.User{}, "", err
	}

	inputPassword := strings.TrimSpace(password)
	storedPassword := strings.TrimSpace(userPtr.Password)

	log.Printf("[AuthService] Comparing passwords: '%s' vs '%s'", inputPassword, storedPassword)

	if storedPassword != inputPassword {
		return model.User{}, "", errors.New("invalid credentials")
	}

	sessionID := util.GenerateSessionID()
	if err := repository.SaveUserSession(userPtr.ID, sessionID); err != nil {
		return model.User{}, "", err
	}

	return *userPtr, sessionID, nil
}
