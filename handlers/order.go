package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JackCloudman/MicroserviciosTest/myTypes"
	"github.com/JackCloudman/MicroserviciosTest/test"
	"github.com/JackCloudman/MicroserviciosTest/utils"
)

// Ordenar es un handler para poder ordenar una Comida
// corresponde a /api/sps/helloworld/v1/ordenar
func Ordenar(w http.ResponseWriter, r *http.Request) {
	orden := myTypes.Orden{}
	err := json.NewDecoder(r.Body).Decode(&orden)
	if err != nil {
		utils.WriteResponse(w, "Error en la estructura del body", http.StatusBadRequest)
		return
	}
	if !test.ExistUsuario(orden.ClienteID) {
		utils.WriteResponse(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}
	if !test.ExistComida(orden.ComidaID) {
		utils.WriteResponse(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}
	fmt.Printf("%+v", orden)
	utils.WriteResponse(w, "Orden realizada con exito!", http.StatusOK)
}
