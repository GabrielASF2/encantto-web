package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnectionString string
	ServerPort         string
}

func LoadConfig() *Config {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado, usando variáveis de ambiente existentes.")
	}

	return &Config{
		DBConnectionString: getEnv("DB_CONNECTION_STRING", "postgres://user:password@localhost:5432/dbname?sslmode=disable"),
		ServerPort:         getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
