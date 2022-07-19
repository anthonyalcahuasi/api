package models

import (
	"fmt"
	"strings"
	"time"
	"gorm.io/gorm"
)

type Recurso struct {
	ID            uint   `gorm:"primary_key;column:id" json:"id"`
	TITULO        string `json:"titulo"`
	DESCRIPCION   string `json:"descripcion"`
	SesionID      string `gorm:"size:191"`
	Sesion        Sesion
	TipoRecursoID string `gorm:"size:191"`
	TipoRecurso   TipoRecurso
	RecursoView []RecursoView
}

type TipoRecurso struct {
	ID      uint   `gorm:"primary_key;column:id" json:"id"`
	NOMBRE  string `json:"nombre"`
	ESTADO  string `json:"estado"`
	Recurso []Recurso
}
type RecursoView struct {
	gorm.Model
	IDUSUARIO  string `json:"id_usuario"` 
	USUARIO  string `json:"usuario"`
	APELLIDOS string `json:"apellidos"`
	FOTO string `json:"foto"`
	RecursoID string `gorm:"size:191"`
	Recurso Recurso
}

type ToolVideoConferencia struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	TIPO        string    `json:"tipo"`
	FECHAINICIO time.Time `json:"fechainicio"`
	FECHAFIN    time.Time `json:"fechafin"`
	ENLACE      string    `json:"enlace"`
	ESTADO      string    `json:"estado"`
	Recurso     Recurso   `gorm:"foreignKey:ID; references:id"`
}

type ToolVideo struct {
	ID      uint    `gorm:"primary_key;column:id" json:"id"`
	TIPO    string  `json:"tipo"`
	ENLACE  string  `json:"enlace"`
	Recurso Recurso `gorm:"foreignKey:ID; references:id"`
}
type ToolEnlace struct {
	ID      uint    `gorm:"primary_key;column:id" json:"id"`
	ENLACE  string  `json:"enlace"`
	Recurso Recurso `gorm:"foreignKey:ID; references:id"`
}

type ToolDocumento struct {
	ID      uint    `gorm:"primary_key;column:id" json:"id"`
	ENLACE  string  `json:"enlace"`
	Recurso Recurso `gorm:"foreignKey:ID; references:id"`
}

type ToolActividad struct {
	ID          uint    `gorm:"primary_key;column:id" json:"id"`
	ADJUNTARDOC string  `json:"adjuntardoc"`
	FECHAINICIO string  `json:"fechainicio"`
	FECHAFIN    string  `json:"fechafin"`
	ESTADO      string  `json:"estado"`
	Recurso     Recurso `gorm:"foreignKey:ID; references:id"`
	Trabajo     []Trabajo
}

type ToolForo struct {
	ID          uint    `gorm:"primary_key;column:id" json:"id"`
	FECHAINICIO string  `json:"fechainicio"`
	FECHAFIN    string  `json:"fechafin"`
	Recurso     Recurso `gorm:"foreignKey:ID; references:id"`
	Foro        []Foro
}

type ToolAnuncio struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	TIPO        string    `json:"fechainicio"`
	FECHAINICIO time.Time `json:"fechainicio"`
	FECHAFIN    time.Time `json:"fechafin"`
	Recurso     Recurso   `gorm:"foreignKey:ID; references:id"`
}
type ToolExamen2 struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	ToolExamenID    string `gorm:"size:191"`
	ToolExamen ToolExamen
	Recurso     Recurso   `gorm:"foreignKey:ID; references:id"`
}



func (r *RecursoView) Register(conn *gorm.DB) error {
	r.IDUSUARIO = strings.ToUpper(r.IDUSUARIO)
	r.USUARIO = strings.ToUpper(r.USUARIO)
	r.APELLIDOS = strings.ToUpper(r.APELLIDOS)
	r.RecursoID = strings.ToUpper(r.RecursoID)
	
	var userLookup RecursoView
	var inscripcion Inscripcion
	var err error
	if err = conn.Where("estudiante_id = ? ",r.IDUSUARIO).First(&inscripcion).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("No está inscrito en el curso")
	}
	if err = conn.Find(&userLookup, "id_usuario = ? AND usuario = ? AND apell_id_os = ? AND recurso_id = ?",r.IDUSUARIO, r.USUARIO, r.APELLIDOS, r.RecursoID).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return fmt.Errorf("Error p" + err.Error())
	}
	
	if r.IDUSUARIO == strings.ToUpper(userLookup.IDUSUARIO) && r.USUARIO == strings.ToUpper(userLookup.USUARIO) && r.APELLIDOS == strings.ToUpper(userLookup.APELLIDOS) && r.RecursoID == strings.ToLower(userLookup.RecursoID) {
		return fmt.Errorf("Ya vio el recurso!")
	}
	conn.Create(&r)
	return err
}