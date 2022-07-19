package models

type Horario struct {
	ID           uint   `gorm:"primary_key;column:id" json:"id"`
	FechaInicio    string `json:"fechainicio"`
	FechaFin       string `json:"fechafin"`
	Fecha          string `json:"fecha"`
	CargaDocenteID string `gorm:"size:191"`
	CargaDocente   CargaDocente
}
