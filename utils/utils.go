package utils

import (
	"encoding/json"
	"net/http"

	"github.com/JackCloudman/MicroserviciosTest/myTypes"
)

// WriteResponse es una funcion que nos ayuda a escribir la Respuesta
// recibe el ResponseWriter, el mensaje en bytes y un status code
func WriteResponse(w http.ResponseWriter, mensaje string, code int) {
	message := myTypes.Respuesta{
		Message: mensaje,
	}
	response, _ := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
