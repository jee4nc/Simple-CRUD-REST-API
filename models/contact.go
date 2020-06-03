package models

import (
	"github.com/jinzhu/gorm"
)

//Contact modelo para contactos
type Contact struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Edad        int    `json:"edad"`
	Telefono    string `json:"telefono" gorm:"size:20"`
	Direccion   string `json:"email"`
	Descripcion string `json:"descripcion" gorm:"type:TEXT"`
	// CreatedAt   time.Time `db:"created_at"`
	// UpdatedAt   time.Time `db:"updated_at"`
}
