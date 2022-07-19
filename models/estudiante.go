package models

type Estudiante struct {
	ID          uint   `gorm:"primary_key;column:id" json:"id"`
	CODIGO      string `json:"codigo"`
	User        User   `gorm:"foreignKey:ID; references:id"`
	Inscripcion []Inscripcion
}
