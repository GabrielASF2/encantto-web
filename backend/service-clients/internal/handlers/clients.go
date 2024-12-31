package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/GabrielASF2/encantto-web/backend/service-clients/internal/models"
	"github.com/GabrielASF2/encantto-web/backend/service-clients/internal/services"

	"github.com/gorilla/mux"
)

func RegisterClientRoutes(router *mux.Router, db *sql.DB) {
	service := services.ClientService{DB: db}

	router.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		clients, err := service.GetAllClients()
		if err != nil {
			http.Error(w, "Erro ao buscar clientes", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(clients)
	}).Methods("GET")

	router.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		var client models.Client
		if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
			http.Error(w, "Dados inválidos", http.StatusBadRequest)
			return
		}
		if err := service.CreateClient(&client); err != nil {
			http.Error(w, "Erro ao criar cliente", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}).Methods("POST")
}
