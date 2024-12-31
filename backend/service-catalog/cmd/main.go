package main

import (
	"log"
	"net/http"

	"github.com/GabrielASF2/encantto-web/backend/service-catalog/internal/handlers"

	"github.com/GabrielASF2/encantto-web/backend/shared/config"
	"github.com/GabrielASF2/encantto-web/backend/shared/db"
	"github.com/GabrielASF2/encantto-web/backend/shared/logger"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.Connect()
	if err != nil {
		logger.Error("Erro ao conectar ao banco de dados: " + err.Error())
		return
	}
	defer database.Close()

	router := mux.NewRouter()
	handlers.RegisterCatalogRoutes(router, database)

	logger.Info("Serviço de Catálogo rodando na porta " + cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
