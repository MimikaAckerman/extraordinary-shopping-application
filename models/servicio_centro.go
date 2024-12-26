package models

import "database/sql"

type ServicioCentro struct {
	Center string         `json:"center"`
	CodNav string         `json:"cod_nav"`
	Zone   sql.NullString `json:"zone"` // Maneja valores NULL de la columna Zone
}
