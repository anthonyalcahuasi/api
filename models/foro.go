package models

import "gorm.io/gorm"

type Foro struct {
	gorm.Model
	ESTADO        string `json:"estado"`
	CONTENIDO     string `json:"contenido"`
	FOROENVIO     string `json:"foroenvio"`
	InscripcionID string `gorm:"size:191"`
	Inscripcion   Inscripcion
	ToolForoID    string `gorm:"size:191"`
	ToolForo      ToolForo
}
