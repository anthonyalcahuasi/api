package models

type Sesion struct {
	ID          uint   `gorm:"primary_key;column:id" json:"id"`
	NUMERO      string `json:"numero"`
	CONTENIDO   string `json:"contenido"`
	APRENDIZAJE string `json:"aprendizaje"`
	ESTRATEGIAS string `json:"estrategias"`
	UnidadID    string `gorm:"size:191"`
	Unidad      Unidad
	Recurso     []Recurso
	ToolExamenDetalleUsuario []ToolExamenDetalleUsuario
}
