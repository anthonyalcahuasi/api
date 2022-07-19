package models

import "gorm.io/gorm"

type Trabajo struct {
	gorm.Model
	NOTA            string `json:"nota"`
	ADJUNTARDOC     string `json:"adjuntardoc"`
	ESTADO          string `json:"estado"`
	TRABAJOENVIO    string `json:"trabajoenvio"`
	InscripcionID   string `gorm:"size:191"`
	Inscripcion     Inscripcion
	ToolActividadID string `gorm:"size:191"`
	ToolActividad   ToolActividad
}
