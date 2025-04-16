package repository

import (
	"cecil-ecommerce/internal/model"
	"cecil-ecommerce/internal/util"
	"database/sql"
	"errors"
)

type UserRepository interface {
	Create(username, email, password string) error
	GetByUsername(username string) (*model.User, error)
	GetBySessionID(sessionID string) (model.User, error)
	SaveSession(userID int, sessionID string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(username, email, password string) error {
	hashed, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	var exists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)", username, email).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")
	}

	_, err = r.db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, hashed)
	return err
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	row := r.db.QueryRow(`
		SELECT id, username, email, password, COALESCE(session_id, '') 
		FROM users 
		WHERE username = $1
	`, username)

	var u model.User
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.SessionID); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) GetBySessionID(sessionID string) (model.User, error) {
	var u model.User
	err := r.db.QueryRow(`
		SELECT id, username, email, password, COALESCE(session_id, '')
		FROM users 
		WHERE session_id = $1
	`, sessionID).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.SessionID)
	return u, err
}

func (r *userRepository) SaveSession(userID int, sessionID string) error {
	_, err := r.db.Exec(`
		UPDATE users SET session_id = $1 WHERE id = $2
	`, sessionID, userID)
	return err
}
