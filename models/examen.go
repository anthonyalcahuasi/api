package models

import (
	"time"
)
type ToolExamen struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	TITULO        string    `json:"titulo"`
	DESCRIPCION      string    `json:"descripcion"`
	ToolExamenCategoriaID    string `gorm:"size:191"`
	ToolExamenPregunta []ToolExamenPregunta
	ToolExamenPublicacion []ToolExamenPublicacion
	ToolExamenCategoria ToolExamenCategoria
	ToolExamenDetalleUsuario []ToolExamenDetalleUsuario
	ToolExamen2 []ToolExamen2
}

type ToolExamenPublicacion struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	FECHAINICIO time.Time `json:"fechainicio"`
	FECHAFIN    time.Time `json:"fechafin"`
	ESTADO      string    `json:"descripcion"`
	PUBLICO  string    `json:"publico"`
	ToolExamenID    string `gorm:"size:191"`
	ToolExamen ToolExamen
	
}
type ToolExamenCategoria struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	DESCRIPCION        string    `json:"categoria"`
	ToolExamen []ToolExamen
}
type ToolExamenPregunta struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	TITULO        string    `json:"titulo"`
	IMG      string    `json:"img"`
	TIME time.Time `json:"time"`
	ToolExamenID    string `gorm:"size:191"`
	ToolExamenAlternativa []ToolExamenAlternativa
}


type ToolExamenAlternativa struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	ALTERNATIVA        string    `json:"alternativa"`
	ESCORRECTO      string    `json:"escorrecto"`
	ToolExamenPreguntaID    string `gorm:"size:191"`
}


type ToolExamenDetalleUsuario struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	FECHAINICIO    time.Time `json:"fechainicio"`
	FECHAFIN    time.Time `json:"fechafin"`
	ESTADO      string    `json:"descripcion"`
	USER      string    `json:"usuario"`
	Puntaje  string    `json:"puntaje"`
	SesionID string `gorm:"size:191"`
	Sesion Sesion
	ToolExamenID string `gorm:"size:191"`
	ToolExamen ToolExamen
	SupervicionID string `gorm:"size:191"`
	Supervicion Supervicion
}


type Supervicion struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	IP        string    `json:"ip"`
	LATITUD      string    `json:"latitud"`
	NAVEGADOR      string    `json:"navegador"`
	SISTOPE      string    `json:"sitope"`
	ToolExamenDetalleUsuario []ToolExamenDetalleUsuario
}


