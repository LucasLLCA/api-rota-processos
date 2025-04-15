package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB é exportado para ser usado fora do pacote
var DB *sql.DB

func Connect() error {
	// Carrega as variáveis do .env
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Erro ao carregar .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORLD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Atribui à variável global DB (sem := !)
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("Erro ao abrir conexão com o banco: %v", err)
	}

	// Testa a conexão
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("Erro ao conectar com o banco: %v", err)
	}

	return nil
}
