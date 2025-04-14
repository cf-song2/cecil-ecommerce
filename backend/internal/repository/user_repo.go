package repository

import (
	"cecil-ecommerce/internal/model"
	"errors"
)

func CreateUser(username, email, password string) error {
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)
	`, username, email).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")
	}

	_, err = db.Exec(`
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
	`, username, email, password)

	return err
}

func GetUserByUsername(username string) (*model.User, error) {
	row := db.QueryRow(`
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

func GetUserBySessionID(sessionID string) (model.User, error) {
	var u model.User
	err := db.QueryRow(`
		SELECT id, username, email, password, COALESCE(session_id, '')
		FROM users 
		WHERE session_id = $1
	`, sessionID).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.SessionID)
	return u, err
}

func SaveUserSession(userID int, sessionID string) error {
	_, err := db.Exec(`
		UPDATE users SET session_id = $1 WHERE id = $2
	`, sessionID, userID)
	return err
}
