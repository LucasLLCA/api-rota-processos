package routes

import (
	"github.com/LucasLLCA/api-rota-processos/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/processo/:numero", handlers.GetProcessoByNumero)

	return r
}
