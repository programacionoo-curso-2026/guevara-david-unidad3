package main

import (
	"PlanificadorEficiente/dataaccess"
	"log"
)

func main() {
	// Inicializar la base de datos
	db := dataaccess.InitDB()
	defer db.Close() // Importante: cerrar la conexión al final
	log.Println("Base de datos inicializada correctamente")
}
