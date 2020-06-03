package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jee4nc/REST/models"
	"github.com/jee4nc/REST/utils"
)

//GetContact obtiene un contacto por su ID
func GetContact(w http.ResponseWriter, r *http.Request) {
	//Estructura vacia donde se guardaran los datos
	contact := models.Contact{}
	//Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	//Conexion a la DB
	db := utils.GetConnection()

	defer db.Close()
	//Consulta a la DB - Select * FROM contact WHERE ID = ?
	db.Find(&contact, id)
	//Se comprueba que exista el registro
	if contact.ID > 0 {
		//Se codifican los datos a formato JSON
		j, _ := json.Marshal(contact)
		//Se envian los datos
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		//Si no existe se envia un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}

//GetContacts busca muchos contactos
func GetContacts(w http.ResponseWriter, r *http.Request) {
	//Slice (array) donde se guardaran los datos
	contacts := []models.Contact{}
	//Conexion a la DB
	db := utils.GetConnection()
	defer db.Close()
	//Consulta a la DB -SELECT * FROM contacts
	db.Find(&contacts)
	//Se codifican los datos a formato json
	j, _ := json.Marshal(contacts)
	// Se envian los datos
	utils.SendResponse(w, http.StatusOK, j)

}

// StoreContact GUARDA UN NUEVO CONTACTO
func StoreContact(w http.ResponseWriter, r *http.Request) {
	// Estructura donde se guardan los datos del body
	contact := models.Contact{}
	//COnexion a la base de datos
	db := utils.GetConnection()
	defer db.Close()
	// Se decodifican los datos del bodu
	//a la estructura contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		// SI hay algun error en los datos devolvera
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	//Se guardan los datos en la DB
	err = db.Create(&contact).Error
	if err != nil {
		//Si hay algun error al guardar devolvera
		//error 500
		fmt.Println(err)
		utils.SendErr(w, http.StatusInternalServerError)
		return
	}
	//Se codifica el nuevo registro y se devuelve
	j, _ := json.Marshal(contact)
	utils.SendResponse(w, http.StatusCreated, j)
}

//UpdateContact modifica los datos de un contacto por su ID
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	//EStructura donde se almacenaran los datos

}
