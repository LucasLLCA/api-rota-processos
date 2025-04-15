package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/LucasLLCA/api-rota-processos/config"
	"github.com/LucasLLCA/api-rota-processos/handlers"
)

func main() {
	// Carrega variáveis de ambiente do .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado. Usando variáveis do sistema.")
	}

	// Conecta ao banco
	err = config.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar com banco: %v", err)
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(config.DB)

	// Cria servidor
	r := gin.Default()

	// Rotas
	r.GET("/processo/:numero", handlers.BuscarProcesso)

	// Inicia servidor
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando em http://localhost:%s", port)
	_ = r.Run(":" + port)
}
