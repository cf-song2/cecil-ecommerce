package repository

import (
	"cecil-ecommerce/internal/model"
)

func GetAllProducts() ([]model.Product, error) {
	rows, err := db.Query("SELECT id, name, description, price, image_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductByID(id int) (*model.Product, error) {
	row := db.QueryRow("SELECT id, name, description, price, image_url FROM products WHERE id = $1", id)
	var p model.Product
	if err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL); err != nil {
		return nil, err
	}
	return &p, nil
}

func CreateProductIfNotExists(p model.Product) error {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE name = $1)", p.Name).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return CreateProduct(p)
}

func CreateProduct(p model.Product) error {
	_, err := db.Exec(`
		INSERT INTO products (name, description, price, image_url)
		VALUES ($1, $2, $3, $4)
	`, p.Name, p.Description, p.Price, p.ImageURL)
	return err
}
