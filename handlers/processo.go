package handlers

import (
	"database/sql"
	"net/http"

	"github.com/LucasLLCA/api-rota-processos/config"
	"github.com/LucasLLCA/api-rota-processos/models"
	"github.com/gin-gonic/gin"
)

func GetProcessoByNumero(c *gin.Context) {
	numero := c.Param("numero")

	query := `
		SELECT protocol_id, protocol, created, type_id, type_name, sector_id
		FROM protocol
		WHERE protocol = $1
	`

	var processo models.Processo
	err := config.DB.QueryRow(query, numero).Scan(
		&processo.ProtocolID,
		&processo.Protocol,
		&processo.Created,
		&processo.TypeID,
		&processo.TypeName,
		&processo.SectorID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Processo não encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar processo"})
		}
		return
	}

	c.JSON(http.StatusOK, processo)
}
