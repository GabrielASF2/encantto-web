package services

import (
	"database/sql"
	"service-clients/internal/models"
)

type ClientService struct {
	DB *sql.DB
}

func (s *ClientService) GetAllClients() ([]models.Client, error) {
	rows, err := s.DB.Query("SELECT id, name, email, phone FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Phone); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func (s *ClientService) CreateClient(client *models.Client) error {
	_, err := s.DB.Exec("INSERT INTO clients (name, email, phone) VALUES ($1, $2, $3)",
		client.Name, client.Email, client.Phone)
	return err
}
