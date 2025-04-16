package service

import (
	"cecil-ecommerce/internal/model"
	"cecil-ecommerce/internal/repository"
	"cecil-ecommerce/internal/util"
	"errors"
	"strings"
)

type AuthService interface {
	Authenticate(username, password string) (model.User, string, error)
	Register(username, email, password string) error
	GetCurrentUser(sessionID string) (model.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(r repository.UserRepository) AuthService {
	return &authService{userRepo: r}
}

func (s *authService) Authenticate(username, password string) (model.User, string, error) {
	userPtr, err := s.userRepo.GetByUsername(strings.TrimSpace(username))
	if err != nil {
		return model.User{}, "", err
	}

	if !util.CheckPasswordHash(strings.TrimSpace(password), userPtr.Password) {
		return model.User{}, "", errors.New("invalid credentials")
	}

	sessionID := util.GenerateSessionID()
	if err := s.userRepo.SaveSession(userPtr.ID, sessionID); err != nil {
		return model.User{}, "", err
	}

	return *userPtr, sessionID, nil
}

func (s *authService) Register(username, email, password string) error {
	return s.userRepo.Create(username, email, password)
}

func (s *authService) GetCurrentUser(sessionID string) (model.User, error) {
	return s.userRepo.GetBySessionID(sessionID)
}
