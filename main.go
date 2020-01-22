package main

import (
	"log"
	"net/http"

	"github.com/JackCloudman/MicroserviciosTest/settings"
)

const (
	host = "0.0.0.0:8090"
)

func main() {
	router := settings.GetRoutes()
	log.Fatal(http.ListenAndServe(host, router))
}
