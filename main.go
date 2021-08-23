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

	// Inicio DB:
	startDb(schemaObject)

	todayUnix := time.Now().Unix() // Default date
	// Tenemos una fecha y traemos la data de las API externas:
	todayData := fetchDayData(todayUnix, queryProducts, queryOrigins)

	addDayData(todayData)

	fmt.Println("UPLOAD finished...")

	allBuyers := getAllBuyers(queryAllBuyers)
	fmt.Println(string(jsoniterMarshall(allBuyers.Q, "dgraph")))

	vars := map[string]string{"$id": "2b85fd40"}
	buyerData := getBuyerById(queryBuyerById, vars)
	fmt.Println(string(jsoniterMarshall(buyerData.Owner, "dgraph")))
	fmt.Println(string(jsoniterMarshall(buyerData.OtherBuyers, "dgraph")))
	fmt.Println(string(jsoniterMarshall(buyerData.OtherProducts, "dgraph")))

}
