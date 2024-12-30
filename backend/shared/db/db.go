package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // Driver para PostgreSQL
)

func Connect() (*sql.DB, error) {
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Configuração de tempo limite
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Testar conexão
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexão com banco de dados bem-sucedida!")
	return db, nil
}
