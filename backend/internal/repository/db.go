package repository

import (
	"database/sql"
	"errors"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func NewRepositories() (*Repositories, error) {
	if db == nil {
		return nil, errors.New("database is not initialized")
	}
	return &Repositories{
		Product: NewProductRepository(db),
		User:    NewUserRepository(db),
		Cart:    NewCartRepository(db),
	}, nil
}
