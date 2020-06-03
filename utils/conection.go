package utils

import (
	"log"

	"github.com/jinzhu/gorm"
	// MYSQL
)

//GetConnection obtiuene una conexion a la base de datos
func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql",
		"userdb:pswdb@/dbname?charset=utf8&parseTime=True&Loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
