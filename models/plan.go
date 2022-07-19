package models

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Nivel struct {
	ID     uint   `gorm:"primary_key;column:id" json:"id"`
	NOMBRE string `json:"nombre"`
	ESTADO string `json:"estado"`
	Grado  []Grado
}

type Grado struct {
	ID        uint   `gorm:"primary_key;column:id" json:"id"`
	NOMBRE    string `json:"nombre"`
	ESTADO    string `json:"estado"`
	NivelID   string `gorm:"size:191"`
	Nivel     Nivel
	Jerarquia []Jerarquia
}

type Sucursal struct {
	ID        uint   `gorm:"primary_key;column:id" json:"id"`
	NOMBRE    string `json:"nombre"`
	ESTADO    string `json:"estado"`
	Jerarquia []Jerarquia
}
type Jerarquia struct {
	ID         uint   `gorm:"primary_key;column:id" json:"id"`
	ESTADO     string `json:"estado"`
	GradoID    string `gorm:"size:191"`
	SucursalID string `gorm:"size:191"`
	Grado      Grado
	Sucursal   Sucursal
	Plan       []Plan
}

type Curso struct {
	ID          uint   `gorm:"primary_key;column:id" json:"id"`
	Nombre      string `json:"nombre"`
	Estado      string `json:"estado"`
	Informacion string `json:"informacion"`
	Sumilla     string `json:"sumilla"`
	Competencia string `json:"competencia"`
	Img         string `json:"img"`
	Color       string `json:"color"`
	Plan        []Plan
}

type Plan struct {
	ID           uint   `gorm:"primary_key;column:id" json:"id"`
	ESTADO       string `json:"estado"`
	JerarquiaID  string `gorm:"size:191"`
	CursoID      string `gorm:"size:191"`
	Curso        Curso
	Jerarquia    Jerarquia
	CargaDocente []CargaDocente
}

func (p *Plan) Register(conn *gorm.DB) error {
	p.JerarquiaID = strings.ToLower(p.JerarquiaID)
	p.CursoID = strings.ToLower(p.CursoID)
	var userLookup Plan
	var err error
	if err = conn.Find(&userLookup, "jerarquia_id = ? AND curso_id = ?", p.JerarquiaID, p.CursoID).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return fmt.Errorf("Error p" + err.Error())
	}
	if p.JerarquiaID == strings.ToLower(userLookup.JerarquiaID) && p.CursoID == strings.ToLower(userLookup.CursoID) {
		return fmt.Errorf("Ya Existe el Plan!")
	}
	conn.Create(&p)
	return err // ya te asigna el ID del user
}

func (j *Jerarquia) Register(conn *gorm.DB) error {
	j.GradoID = strings.ToLower(j.GradoID)
	j.SucursalID = strings.ToLower(j.SucursalID)
	var jerarquiaLookup Jerarquia
	var err error
	if err = conn.Find(&jerarquiaLookup, "grado_id = ? AND sucursal_id = ?", j.GradoID, j.SucursalID).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return fmt.Errorf("Error p" + err.Error())
	}
	if j.GradoID == strings.ToLower(jerarquiaLookup.GradoID) && j.SucursalID == strings.ToLower(jerarquiaLookup.SucursalID) {
		return fmt.Errorf("Existe la Jerárquia!")
	}
	conn.Create(&j)
	return err // ya te asigna el ID del user
}
