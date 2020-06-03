package routes

import (
	"github.com/gorilla/mux"
	"github.com/jee4nc/REST/controllers"
)

//SetContactsRoutes agrega las rutas de contactos

func SetContactsRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()

	subRouter.HandleFunc("/contacts/{id}",
		controllers.GetContact).Methods("GET")

	subRouter.HandleFunc("/contacts",
		controllers.GetContacts).Methods("GET")

	subRouter.HandleFunc("/contacts/{id}",
		controllers.StoreContact).Methods("POST")

	subRouter.HandleFunc("/contacts/{id}",
		controllers.UpdateContact).Methods("PUT")

	subRouter.HandleFunc("/contacts/{id}",
		controllers.DeleteContact).Methods("DELETE")

}
