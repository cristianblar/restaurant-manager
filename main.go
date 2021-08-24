package main

import (
	"github.com/joho/godotenv"
)

func main() {

	// Carga de variables de entorno
	dotenvError := godotenv.Load()
	panicErrorHandler(dotenvError)

	server := CreateServer()
	server.Listen()

}
