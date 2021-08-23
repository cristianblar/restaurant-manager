package main

var dbInstance *DatabaseConnection = nil

func startDb(schema string) {
	dbInstance = CreateDatabase(schema)
}

func prepareNewDate() {
	dbInstance.DropData()
}

func initialProductsOperation(query string, productsToDbChannel <-chan []*Product, productsFromDbChannel chan<- *ProductQuery) {

	productsToDb := <-productsToDbChannel

	productsJson := jsoniterMarshall(productsToDb, "dgraph")
	dbInstance.BulkJsonMutation(productsJson)

	queryJson := dbInstance.GetQuery(query)

	queryResult := new(ProductQuery)
	jsoniterUnmarshall(queryJson, queryResult, "dgraph")

	productsFromDbChannel <- queryResult
	close(productsFromDbChannel)

}

func initialOriginsOperation(query string, originsToDbChannel <-chan []*Origin, originsFromDbChannel chan<- *OriginQuery) {

	originsToDb := <-originsToDbChannel

	originsJson := jsoniterMarshall(originsToDb, "dgraph")
	dbInstance.BulkJsonMutation(originsJson)

	queryJson := dbInstance.GetQuery(query)

	queryResult := new(OriginQuery)
	jsoniterUnmarshall(queryJson, queryResult, "dgraph")

	originsFromDbChannel <- queryResult
	close(originsFromDbChannel)

}

func addDayData(dayData []*Buyer) {

	dayDataJson := jsoniterMarshall(dayData, "dgraph")
	dbInstance.BulkJsonMutation(dayDataJson)

}

func getAllBuyers(query string) *AllBuyersQuery {

	queryJson := dbInstance.GetQuery(query)

	queryResult := new(AllBuyersQuery)
	jsoniterUnmarshall(queryJson, queryResult, "dgraph")

	return queryResult

}

func getBuyerById(query string, vars map[string]string) *BuyerQuery {

	queryJson := dbInstance.GetQueryWithVariables(query, vars)

	queryResult := new(BuyerQuery)
	jsoniterUnmarshall(queryJson, queryResult, "dgraph")

	return queryResult

}
