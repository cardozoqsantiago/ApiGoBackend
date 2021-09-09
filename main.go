package main

import (
	"log"

	"github.com/cardozoqsantiago/ApiGoBackend/bd"
	"github.com/cardozoqsantiago/ApiGoBackend/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la bd")
		return
	}
	handlers.Manejadores()

}
