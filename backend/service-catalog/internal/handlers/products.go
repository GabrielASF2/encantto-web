package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/GabrielASF2/encantto-web/backend/service-catalog/internal/models"
	"github.com/GabrielASF2/encantto-web/backend/service-catalog/internal/services"

	"github.com/gorilla/mux"
)

func RegisterCatalogRoutes(router *mux.Router, db *sql.DB) {
	service := services.ProductService{DB: db}

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products, err := service.GetAllProducts()
		if err != nil {
			http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(products)
	}).Methods("GET")

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		var product models.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Dados inválidos", http.StatusBadRequest)
			return
		}
		if err := service.CreateProduct(&product); err != nil {
			http.Error(w, "Erro ao criar produto", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}).Methods("POST")
}
