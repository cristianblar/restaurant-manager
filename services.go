package main

var dbInstance *DatabaseConnection = nil

func startDb(schema string) {
	dbInstance = CreateDatabase(schema)
}

func prepareNewDate() error {
	return dbInstance.DropData()
}

// Para crear cada producto una única vez en la DB:
func initialProductsOperation(query string, productsToDbChannel <-chan []*Product, productsFromDbChannel chan<- *ProductQuery) {

	productsToDb := <-productsToDbChannel

	productsJson, marshallError := jsoniterMarshall(productsToDb, "dgraph")
	panicErrorHandler(marshallError)

	dbInstance.BulkJsonMutation(productsJson)

	queryJson, queryError := dbInstance.GetQuery(query)
	panicErrorHandler(queryError)

	queryResult := new(ProductQuery)
	unmarshallError := jsoniterUnmarshall(queryJson, queryResult, "dgraph")
	panicErrorHandler(unmarshallError)

	productsFromDbChannel <- queryResult
	close(productsFromDbChannel)

}

// Para crear cada origen (device + ip) una única vez en la DB:
func initialOriginsOperation(query string, originsToDbChannel <-chan []*Origin, originsFromDbChannel chan<- *OriginQuery) {

	originsToDb := <-originsToDbChannel

	originsJson, marshallError := jsoniterMarshall(originsToDb, "dgraph")
	panicErrorHandler(marshallError)

	dbInstance.BulkJsonMutation(originsJson)

	queryJson, queryError := dbInstance.GetQuery(query)
	panicErrorHandler(queryError)

	queryResult := new(OriginQuery)
	unmarshallError := jsoniterUnmarshall(queryJson, queryResult, "dgraph")
	panicErrorHandler(unmarshallError)

	originsFromDbChannel <- queryResult
	close(originsFromDbChannel)

}

func addDayData(dayData []*Buyer) {

	dayDataJson, marshallError := jsoniterMarshall(dayData, "dgraph")
	panicErrorHandler(marshallError)

	dbInstance.BulkJsonMutation(dayDataJson)

}

func getAllBuyers(query string) ([]byte, *AllBuyersQuery, error) {

	queryJson, queryError := dbInstance.GetQuery(query)
	if queryError != nil {
		return nil, nil, queryError
	}

	queryResult := new(AllBuyersQuery)
	unmarshallError := jsoniterUnmarshall(queryJson, queryResult, "dgraph")
	if unmarshallError != nil {
		return nil, nil, unmarshallError
	}

	// Fecha sincronizada no tiene compradores...
	if len(queryResult.Q) == 0 {
		return nil, nil, nil
	}

	return queryJson, queryResult, nil

}

func getBuyerById(query string, vars map[string]string) ([]byte, *BuyerQuery, error) {

	queryJson, queryError := dbInstance.GetQueryWithVariables(query, vars)
	if queryError != nil {
		return nil, nil, queryError
	}

	queryResult := new(BuyerQuery)
	unmarshallError := jsoniterUnmarshall(queryJson, queryResult, "dgraph")
	if unmarshallError != nil {
		return nil, nil, unmarshallError
	}

	// ID no existe...
	if len(queryResult.Owner) == 0 {
		return nil, nil, nil
	}

	return queryJson, queryResult, nil

}
