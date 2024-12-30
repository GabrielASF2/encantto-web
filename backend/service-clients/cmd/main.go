package main

import (
	"log"
	"net/http"

	"service-clients/internal/handlers"

	"github.com/GabrielASF2/encantto-web/shared/config"
	"github.com/GabrielASF2/encantto-web/shared/db"
	"github.com/GabrielASF2/encantto-web/shared/logger"
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
	handlers.RegisterClientRoutes(router, database)

	logger.Info("Serviço de Clientes rodando na porta " + cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
