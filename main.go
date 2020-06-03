package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jee4nc/REST/routes"
	"github.com/jee4nc/REST/utils"
)

func main() {
	//Migracion de la base de datos
	utils.MigrateDB()
	//router para el manejo de las rutas de la aplicacion
	r := mux.NewRouter()
	//Se agregan las rutas de contactos
	routes.SetContactsRoutes(r)
	//Generacioin de un nuevo servidor, especificamos
	// puerto y las rutas

	srv := http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	log.Println("Running on port 4000")
	//Se ejecuta en el servidor
	log.Println(srv.ListenAndServe())
}
