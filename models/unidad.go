package models

type Unidad struct {
	ID             uint   `gorm:"primary_key;column:id" json:"id"`
	NUMERO         string `json:"numero"`
	TEMA           string `json:"tema"`
	CargaDocenteID string `gorm:"size:191"`
	Sesion         []Sesion
	CargaDocente   CargaDocente
}
