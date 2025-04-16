package config

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	// Carrega .env
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("erro ao carregar .env: %v", err)
	}

	// Configura a string de conexão
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"10.0.122.89",
		"5432",
		"painel",
		"Jw{d1*r!.q}FwZV",
		"painel_sead_hml",
	)
	fmt.Println("DSN:", connStr)

	// Conecta ao banco
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("erro ao abrir conexão: %v", err)
	}

	// Testa a conexão
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("erro ao pingar o banco: %v", err)
	}

	fmt.Println("Conectado ao PostgreSQL!")
	return nil
}
