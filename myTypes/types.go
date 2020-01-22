package myTypes

import "time"

// Usuario estructura que contiene la informacion del usuario
type Usuario struct {
	ID       uint64
	Nombre   string
	Apellido string
}

// Restaurante es la estrucutra que contiene la informacion de un Restaurante
type Restaurante struct {
	ID        uint64
	Nombre    string
	Direccion string
}

// Comida Información de los platillos
type Comida struct {
	ID          uint64
	Nombre      string
	Descripcion string
	Restaurante uint64
}

// Orden estructura que contiene los datos de las ordenes
type Orden struct {
	ID        uint64    `json:"-"`
	ClienteID uint64    `json:"id_cliente"`
	ComidaID  uint64    `json:"id_comida"`
	OrderTime time.Time `json:"time"`
}

// Respuesta es la estructura que contiene la respuesta
type Respuesta struct {
	Message string `json:"message"`
}
