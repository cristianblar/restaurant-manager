package lib

import (
	"github.com/cristianblar/restaurant-manager/api/database"
	"github.com/cristianblar/restaurant-manager/api/utils"
)

var dbInstance *database.DatabaseConnection = nil

func StartDb(schema string) {
	dbInstance = database.CreateDatabase(schema)
}

func PrepareNewDate() error {
	return dbInstance.DropData()
}

// Para crear cada producto una única vez en la DB:
func initialProductsOperation(query string, productsToDbChannel <-chan []*Product, productsFromDbChannel chan<- *ProductQuery) {

	productsToDb := <-productsToDbChannel

	productsJson, marshallError := utils.JsoniterMarshall(productsToDb, "dgraph")
	utils.PanicErrorHandler(marshallError)

	dbInstance.BulkJsonMutation(productsJson)

	queryJson, queryError := dbInstance.GetQuery(query)
	utils.PanicErrorHandler(queryError)

	queryResult := new(ProductQuery)
	unmarshallError := utils.JsoniterUnmarshall(queryJson, queryResult, "dgraph")
	utils.PanicErrorHandler(unmarshallError)

	productsFromDbChannel <- queryResult
	close(productsFromDbChannel)

}

// Para crear cada origen (device + ip) una única vez en la DB:
func initialOriginsOperation(query string, originsToDbChannel <-chan []*Origin, originsFromDbChannel chan<- *OriginQuery) {

	originsToDb := <-originsToDbChannel

	originsJson, marshallError := utils.JsoniterMarshall(originsToDb, "dgraph")
	utils.PanicErrorHandler(marshallError)

	dbInstance.BulkJsonMutation(originsJson)

	queryJson, queryError := dbInstance.GetQuery(query)
	utils.PanicErrorHandler(queryError)

	queryResult := new(OriginQuery)
	unmarshallError := utils.JsoniterUnmarshall(queryJson, queryResult, "dgraph")
	utils.PanicErrorHandler(unmarshallError)

	originsFromDbChannel <- queryResult
	close(originsFromDbChannel)

}

func AddDayData(dayData []*Buyer) {

	dayDataJson, marshallError := utils.JsoniterMarshall(dayData, "dgraph")
	utils.PanicErrorHandler(marshallError)

	dbInstance.BulkJsonMutation(dayDataJson)

}

func GetAllBuyers(query string) ([]byte, *AllBuyersQuery, error) {

	queryJson, queryError := dbInstance.GetQuery(query)
	if queryError != nil {
		return nil, nil, queryError
	}

	queryResult := new(AllBuyersQuery)
	unmarshallError := utils.JsoniterUnmarshall(queryJson, queryResult, "dgraph")
	if unmarshallError != nil {
		return nil, nil, unmarshallError
	}

	// Fecha sincronizada no tiene compradores...
	if len(queryResult.Q) == 0 {
		return nil, nil, nil
	}

	return queryJson, queryResult, nil

}

func GetBuyerById(query string, vars map[string]string) ([]byte, *BuyerQuery, error) {

	queryJson, queryError := dbInstance.GetQueryWithVariables(query, vars)
	if queryError != nil {
		return nil, nil, queryError
	}

	queryResult := new(BuyerQuery)
	unmarshallError := utils.JsoniterUnmarshall(queryJson, queryResult, "dgraph")
	if unmarshallError != nil {
		return nil, nil, unmarshallError
	}

	// ID no existe...
	if len(queryResult.Owner) == 0 {
		return nil, nil, nil
	}

	return queryJson, queryResult, nil

}
