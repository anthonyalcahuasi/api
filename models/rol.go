package models

type Rol struct {
	ID     uint   `gorm:"primary_key;column:id" json:"id"`
	NOMBRE string `json:"nombre"`
	ESTADO string `json:"estado"`
}

type UserRol struct {
	UserID string `gorm:"column:user_id" json:"user_id"`
	RolID  string `gorm:"column:rol_id" json:"rol_id"`
}

func (UserRol) TableName() string {
	return "user_rol"
}
