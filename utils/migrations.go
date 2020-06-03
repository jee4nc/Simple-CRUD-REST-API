package utils

import (
	"fmt"

	"github.com/jee4nc/REST/models"
)

//MIGRATEDB migra la base de datos
func MigrateDB() {
	db := GetConnection()
	defer db.Close()
	fmt.Println("MIgrating models..")

	//Automigrate se encarga de migrar la base de datos si no se ha migrado
	// y lo hace a partir del modelo
	db.AutoMigrate(&models.Contact{})
}
