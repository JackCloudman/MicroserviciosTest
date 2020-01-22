package test

import (
	. "github.com/JackCloudman/MicroserviciosTest/myTypes"
)

var (
	numOrder = 0
	// Usuarios es una variable que simula una DB con usuarios
	Usuarios = []Usuario{
		Usuario{
			ID:       1,
			Nombre:   "Jack",
			Apellido: "Cloudman",
		},
		Usuario{
			ID:       2,
			Nombre:   "SPS",
			Apellido: "1234",
		},
	}
	// Restaurantes simula una DB con restaurantes
	Restaurantes = []Restaurante{
		Restaurante{
			ID:        1,
			Nombre:    "La fondita 1",
			Direccion: "Calle 1 #2",
		},
	}
	// Platillos simula una DB con comidas
	Platillos = []Comida{
		Comida{
			ID:          1,
			Nombre:      "Sushi",
			Descripcion: "Rollo de sushi 1",
			Restaurante: 1,
		},
		Comida{
			ID:          2,
			Nombre:      "Mole con pollo",
			Descripcion: "Muy rico mole con pollo",
			Restaurante: 1,
		},
	}
)

// ExistUsuario simula "select 1 from usuarios where id=?"
func ExistUsuario(id uint64) bool {
	for _, u := range Usuarios {
		if u.ID == id {
			return true
		}
	}
	return false
}

// ExistComida simula "select 1 from usuarios where id=?"
func ExistComida(id uint64) bool {
	for _, c := range Platillos {
		if c.ID == id {
			return true
		}
	}
	return false
}
