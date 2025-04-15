package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() error {
	// Carrega variáveis do .env
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Erro ao carregar .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// DSN para MySQL: usuário:senha@tcp(host:porta)/banco
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, dbname)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("Erro ao abrir conexão com o banco MySQL: %v", err)
	}

	// Testa a conexão
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("Erro ao conectar com o banco MySQL: %v", err)
	}

	return nil
}
