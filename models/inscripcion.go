package models

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Inscripcion struct {
	gorm.Model
	CargaDocenteID string `gorm:"size:191"`
	CargaDocente   CargaDocente
	EstudianteID   string `gorm:"size:191"`
	Estudiante     Estudiante
}

func (u *Inscripcion) Register(conn *gorm.DB) error {
	u.CargaDocenteID = strings.ToLower(u.CargaDocenteID)
	u.EstudianteID = strings.ToLower(u.EstudianteID)
	var userLookup Inscripcion
	var err error
	if err = conn.Find(&userLookup, "carga_docente_id = ? AND estudiante_id = ?", u.CargaDocenteID, u.EstudianteID).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return fmt.Errorf("Error p" + err.Error())
	}
	if u.CargaDocenteID == strings.ToLower(userLookup.CargaDocenteID) && u.EstudianteID == strings.ToLower(userLookup.EstudianteID) {
		return fmt.Errorf("Ya estás en el curso")
	}
	conn.Create(&u)
	return err // ya te asigna el ID del user
}
