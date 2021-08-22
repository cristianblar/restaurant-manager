package main

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	// Carga de variables de entorno
	dotenvError := godotenv.Load()
	errorHandler(dotenvError)

	// Primera conexión -> Carga Schema:
	connectToDb(true, false, nil, nil)

	todayUnix := time.Now().Unix() // Default date
	// Tenemos una fecha y traemos la data de las API externas:
	todayData := fetchDayData(todayUnix)
	// Data lista para enviar:
	jsonForDb := jsoniterMarshall(todayData, "dgraph")

	fmt.Println(string(jsonForDb))

	// // Envío de la data -> Mutation:
	// connectToDb(false, true, applyMutation, jsonForDb)

	// fmt.Println("UPLOAD finished...")

	// Consulta de la data -> Query:
	// queryJson := connectToDb(false, false, getQuery, []byte(queryAllBuyers))

	// var queryResults struct {
	// 	Q []Buyer
	// }

	// jsoniterUnmarshall(queryJson, &queryResults, "dgraph")

	// for idx, buyer := range queryResults.Q {
	// 	fmt.Printf("%d. Id: %s, Name: %s, Age: %d\n", idx+1, buyer.Id, buyer.Name, buyer.Age)
	// }

}
