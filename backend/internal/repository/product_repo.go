package repository

import (
	"cecil-ecommerce/internal/model"
	"database/sql"
)

type ProductRepository interface {
	GetAll() ([]model.Product, error)
	GetByID(id int) (*model.Product, error)
	Create(p model.Product) error
	CreateIfNotExists(p model.Product) error
}

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) ProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) GetAll() ([]model.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, image_url FROM products")
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

func (r *PostgresProductRepository) GetByID(id int) (*model.Product, error) {
	row := r.db.QueryRow("SELECT id, name, description, price, image_url FROM products WHERE id = $1", id)
	var p model.Product
	if err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PostgresProductRepository) CreateIfNotExists(p model.Product) error {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE name = $1)", p.Name).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return r.Create(p)
}

func (r *PostgresProductRepository) Create(p model.Product) error {
	_, err := r.db.Exec(`
		INSERT INTO products (name, description, price, image_url)
		VALUES ($1, $2, $3, $4)
	`, p.Name, p.Description, p.Price, p.ImageURL)
	return err
}