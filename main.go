package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/LucasLLCA/api-rota-processos/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conecta ao PostgreSQL
	if err := config.Connect(); err != nil {
		log.Fatalf("Falha ao conectar ao banco: %v", err)
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(config.DB)

	// Configura o router
	router := gin.Default()

	// Rota de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Inicia o servidor
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor rodando na porta %s", port)
	_ = router.Run(":" + port)
}
