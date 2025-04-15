package routes

import (
	"api-rota-processos/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/processo/:numero", handlers.GetProcessoByNumero)

	return r
}
