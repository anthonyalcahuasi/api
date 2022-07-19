package models

import (
	"fmt"
	"strings"

	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
)

func (tab *CargaDocente) BeforeCreate(*gorm.DB) error {
	//uuidx := uuid.NewV4()
	tab.CODIGO = randstr.String(10)
	return nil
}

type CargaDocente struct {
	ID          uint   `gorm:"primary_key;column:id" json:"id"`
	DocenteID   string `gorm:"size:191"`
	PeriodoID   string `gorm:"size:191"`
	GrupoID     string `gorm:"size:191"`
	PlanID      string `gorm:"size:191"`
	CODIGO      string `json:"codigo" gorm:"uniqueIndex"`
	Docente     Docente
	Periodo     Periodo
	Grupo       Grupo
	Plan        Plan
	Horario     []Horario
	Unidad      []Unidad
	Inscripcion []Inscripcion
}
type Periodo struct {
	ID           uint   `gorm:"primary_key;column:id" json:"id"`
	NOMBRE       string `json:"nombre"`
	ESTADO       string `json:"estado"`
	CargaDocente []CargaDocente
}
type Grupo struct {
	ID           uint   `gorm:"primary_key;column:id" json:"id"`
	NOMBRE       string `json:"nombre"`
	ESTADO       string `json:"estado"`
	ALIAS        string `json:"alias"`
	CargaDocente []CargaDocente
}

func (g *Grupo) Register(conn *gorm.DB) error {
	g.NOMBRE = strings.ToLower(g.NOMBRE)
	g.ALIAS = strings.ToLower(g.ALIAS)
	var grupoLookup Grupo
	var err error
	if err = conn.Find(&grupoLookup, "nombre = ? AND alias = ?", g.NOMBRE, g.ALIAS).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return fmt.Errorf("Error p" + err.Error())
	}
	if g.NOMBRE == strings.ToLower(grupoLookup.NOMBRE) && g.ALIAS == strings.ToLower(grupoLookup.ALIAS) {
		return fmt.Errorf("Ya Existe el Plan!")
	}
	conn.Create(&g)
	return err // ya te asigna el ID del user
}
