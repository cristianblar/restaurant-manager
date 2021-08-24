package main

import (
	"github.com/cristianblar/restaurant-manager/api/server"
	"github.com/cristianblar/restaurant-manager/api/utils"
	"github.com/joho/godotenv"
)

func main() {

	// Carga de variables de entorno
	dotenvError := godotenv.Load()
	utils.PanicErrorHandler(dotenvError)

	server := server.CreateServer()
	server.Listen()

}
