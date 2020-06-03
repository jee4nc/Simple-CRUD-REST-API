package utils

import (
	"log"

	"github.com/jinzhu/gorm"
	// MYSQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
