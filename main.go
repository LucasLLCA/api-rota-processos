package main

import (
	"log"

	"github.com/LucasLLCA/api-rota-processos/config"
	"github.com/LucasLLCA/api-rota-processos/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.Connect(); err != nil {
		log.Fatalf("Erro ao conectar com banco: %v", err)
	}

	router := gin.Default()

	router.GET("/processo/:numero", handlers.GetProcessoByNumero)

	router.Run(":8080")
}
