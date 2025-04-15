package models

type Processo struct {
	ID       int    `json:"id"`
	Numero   string `json:"numero"`
	CriadoEm string `json:"criado_em"`
	TipoID   int    `json:"tipo_id"`
	TipoNome string `json:"tipo_nome"`
	SetorID  int    `json:"setor_id"`
}
