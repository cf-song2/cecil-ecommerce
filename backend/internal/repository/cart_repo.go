package repository

import (
	"database/sql"
	"cecil-ecommerce/internal/model"
)

type CartRepository interface {
	Add(userID, productID, quantity int) error
	Get(userID int) ([]model.CartItem, error)
	Update(userID, productID, quantity int) error
	Remove(userID, productID int) error
	Clear(userID int) error
}

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) Add(userID, productID, quantity int) error {
	_, err := r.db.Exec(`
		INSERT INTO cart (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, product_id) DO UPDATE
		SET quantity = cart.quantity + EXCLUDED.quantity
	`, userID, productID, quantity)
	return err
}

func (r *cartRepository) Get(userID int) ([]model.CartItem, error) {
	rows, err := r.db.Query(`
		SELECT user_id, product_id, quantity FROM cart WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.CartItem
	for rows.Next() {
		var item model.CartItem
		if err := rows.Scan(&item.UserID, &item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *cartRepository) Update(userID, productID, quantity int) error {
	_, err := r.db.Exec(`
		UPDATE cart SET quantity = $3 WHERE user_id = $1 AND product_id = $2
	`, userID, productID, quantity)
	return err
}

func (r *cartRepository) Remove(userID, productID int) error {
	_, err := r.db.Exec(`
		DELETE FROM cart WHERE user_id = $1 AND product_id = $2
	`, userID, productID)
	return err
}

func (r *cartRepository) Clear(userID int) error {
	_, err := r.db.Exec(`
		DELETE FROM cart WHERE user_id = $1
	`, userID)
	return err
}
