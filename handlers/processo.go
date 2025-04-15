package handlers

import (
	"api-rota-processos/config"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Processo struct {
	ProtocolID int            `json:"protocol_id"`
	Protocol   string         `json:"protocol"`
	Created    sql.NullString `json:"created"`
	TypeID     int            `json:"type_id"`
	TypeName   string         `json:"type_name"`
	SectorID   int            `json:"sector_id"`
}

func GetProcessoByNumero(c *gin.Context) {
	numero := c.Param("numero")

	query := `
		SELECT protocol_id, protocol, created, type_id, type_name, sector_id
		FROM protocol
		WHERE protocol = $1
	`

	var processo Processo
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, processo)
}
