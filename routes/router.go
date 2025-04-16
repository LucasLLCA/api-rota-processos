package routes

import (
	"net/http"

	"github.com/LucasLLCA/api-rota-processos/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Middlewares globais
	r.Use(gin.Logger())   // Log das requisições
	r.Use(gin.Recovery()) // Recupera de panics

	// Rotas
	r.GET("/processo/:numero", handlers.GetProcessoByNumero)

	// Rota de health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
