package models

import "database/sql"

type Processo struct {
	ProtocolID int            `json:"protocol_id"`
	Protocol   string         `json:"protocol"`
	Created    sql.NullString `json:"created"`
	TypeID     int            `json:"type_id"`
	TypeName   string         `json:"type_name"`
	SectorID   int            `json:"sector_id"`
}
