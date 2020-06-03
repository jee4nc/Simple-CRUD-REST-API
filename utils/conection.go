package utils

import (
	"log"

	"github.com/jinzhu/gorm"
	// MYSQL
	_ "github.com/go-sql-driver/mysql"
)

//GetConnection obtiuene una conexion a la base de datos
func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql",
		"root:Jean/0206@/contactos?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
