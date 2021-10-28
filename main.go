package main

import (
	"log"

	"github.com/Juanantogg/server-go-twittor/db"
	"github.com/Juanantogg/server-go-twittor/handlers"
)

func main() {
	if db.ConnectionCheck() {
		log.Fatal("Sin conexion a la Base de Datos")

		return
	}

	handlers.Handlers()
}
