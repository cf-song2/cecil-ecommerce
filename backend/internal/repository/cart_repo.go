package repository

import (
	"cecil-ecommerce/internal/model"
)

func AddToCart(userID, productID, quantity int) error {
	_, err := db.Exec(`
		INSERT INTO cart (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, product_id) DO UPDATE
		SET quantity = cart.quantity + EXCLUDED.quantity
	`, userID, productID, quantity)
	return err
}

func GetCartItems(userID int) ([]model.CartItem, error) {
	rows, err := db.Query(`
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

func UpdateCartItem(userID, productID, quantity int) error {
	_, err := db.Exec(`
		UPDATE cart SET quantity = $3 WHERE user_id = $1 AND product_id = $2
	`, userID, productID, quantity)
	return err
}

func RemoveCartItem(userID, productID int) error {
	_, err := db.Exec(`
		DELETE FROM cart WHERE user_id = $1 AND product_id = $2
	`, userID, productID)
	return err
}

func ClearCart(userID int) error {
	_, err := db.Exec(`
		DELETE FROM cart WHERE user_id = $1
	`, userID)
	return err
}
