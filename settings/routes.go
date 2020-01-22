package settings

import (
	"github.com/JackCloudman/MicroserviciosTest/handlers"
	"github.com/gorilla/mux"
)

// GetRoutes devuelve crea una instancia de mux.Rotuer para
// ayudarnos a manejar los endpoints
func GetRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/sps/helloworld/v1/ordenar", handlers.Ordenar)
	return router
}
