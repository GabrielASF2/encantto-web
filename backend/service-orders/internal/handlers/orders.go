package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"service-orders/internal/models"
	"service-orders/internal/services"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router, db *sql.DB) {
	service := services.OrderService{DB: db}

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		orders, err := service.GetAllOrders()
		if err != nil {
			http.Error(w, "Erro ao buscar pedidos", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}).Methods("GET")

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		var order models.Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, "Dados inválidos", http.StatusBadRequest)
			return
		}
		if err := service.CreateOrder(&order); err != nil {
			http.Error(w, "Erro ao criar pedido", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}).Methods("POST")
}
