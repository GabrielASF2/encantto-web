package services

import (
	"database/sql"
	"encoding/json"

	"github.com/GabrielASF2/encantto-web/backend/service-orders/internal/models"
)

type OrderService struct {
	DB *sql.DB
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	rows, err := s.DB.Query("SELECT id, client_id, product_ids, total FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		var productIDs string

		if err := rows.Scan(&order.ID, &order.ClientID, &productIDs, &order.Total); err != nil {
			return nil, err
		}

		// Convertendo JSON de IDs de produtos para array
		if err := json.Unmarshal([]byte(productIDs), &order.ProductIDs); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	productIDs, err := json.Marshal(order.ProductIDs)
	if err != nil {
		return err
	}

	_, err = s.DB.Exec("INSERT INTO orders (client_id, product_ids, total) VALUES ($1, $2, $3)",
		order.ClientID, productIDs, order.Total)
	return err
}
