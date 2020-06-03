package utils

import "net/http"

//SendResponse envia datos json al cliente
func SendResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

//SendErr envia un error al cliente, con contenido vacio
func SendErr(w http.ResponseWriter, status int) {
	data := []byte(`{}`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
