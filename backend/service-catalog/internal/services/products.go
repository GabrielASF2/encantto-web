package services

import (
	"database/sql"

	"github.com/GabrielASF2/encantto-web/backend/service-catalog/internal/models"
)

type ProductService struct {
	DB *sql.DB
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	rows, err := s.DB.Query("SELECT id, name, price, quantity FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	_, err := s.DB.Exec("INSERT INTO products (name, price, quantity) VALUES ($1, $2, $3)",
		product.Name, product.Price, product.Quantity)
	return err
}
