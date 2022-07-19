package models

type Docente struct {
	ID           uint   `gorm:"primary_key;column:id" json:"id"`
	CODIGO       string `json:"codigo"`
	User         User   `gorm:"foreignKey:ID; references:id"`
	CargaDocente []CargaDocente
}
