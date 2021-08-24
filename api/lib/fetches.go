package lib

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"sync"

	"github.com/cristianblar/restaurant-manager/api/utils"
)

func fetchProducts(date int64, productsChannel chan<- map[string]*Product, productsToDbChannel chan<- []*Product) {

	responseProducts := utils.GenericFetch("products", date)
	defer responseProducts.Body.Close()

	// Creación de reader con csv para leer csv
	csvReader := csv.NewReader(responseProducts.Body)
	// Configuración de csvReader para indicar separador
	csvReader.Comma = '\''
	// Uso de ReadAll para leer todo el csv (se necesita todo)
	productsData, errProductsData := csvReader.ReadAll()
	utils.PanicErrorHandler(errProductsData)

	productsDataProcessed := make(map[string]*Product)
	var productsToDb []*Product

	for _, product := range productsData {
		// Conversión del precio del producto a int
		numPrice, errNumPrice := strconv.Atoi(product[2])
		utils.PanicErrorHandler(errNumPrice)
		// Conversión del precio del producto en centavos a dólares
		var dollarPrice float32 = float32(numPrice) / 100
		newProduct := new(Product)
		newProduct.Id = product[0]
		newProduct.Name = strings.ReplaceAll(product[1], "&", "and")
		newProduct.Price = dollarPrice
		newProduct.DType = "Product"
		productsDataProcessed[newProduct.Id] = newProduct
		productsToDb = append(productsToDb, newProduct)
	}

	productsToDbChannel <- productsToDb
	close(productsToDbChannel)
	productsChannel <- productsDataProcessed
	productsChannel <- productsDataProcessed
	close(productsChannel)

}

func fetchTransactions(date int64, productsChannel <-chan map[string]*Product, transactionsChannel chan<- map[string][]*Transaction, originsChannel chan<- map[string]*Origin, originsToDbChannel chan<- []*Origin) {

	responseTransactions := utils.GenericFetch("transactions", date)
	defer responseTransactions.Body.Close()

	// Creación de reader con bufio para leer por bytes
	transactionsReader := bufio.NewReader(responseTransactions.Body)

	// Creación de mapas para structs
	transactionsDataProcessed := make(map[string][]*Transaction)
	transactionsMap := make(map[string]*Transaction)
	originsMap := make(map[string]*Origin)
	var originsToDb []*Origin

	productsMap := <-productsChannel

	// for forever para la cantidad de transacciones que es desconocida
	for {
		// flags para detener o continuar for forever desde for interno
		loopFlag := false
		emptyFlag := false

		newTransaction := new(Transaction)
		var mapKey string

		// for de 5 iteraciones para leer y guardar cada uno de los 5 datos por transacción
		for i := 0; i < 5; i++ {
			// Lectura de []byte hasta encontrar null -> byte(0) (dato a dato)
			transactionBytes, errTransactionBytes := transactionsReader.ReadBytes(0)

			// Custom error handler para romper for loops al EOF
			if errTransactionBytes != nil {
				if errTransactionBytes == io.EOF {
					loopFlag = true
					break
				} else {
					utils.PanicErrorHandler(errTransactionBytes)
				}
			}

			// Skip de for loops cuando el dato leído es null (fin de transacción)
			if len(transactionBytes) == 1 && transactionBytes[0] == 0 {
				emptyFlag = true
				break
			}

			// Asignación de cada []byte al field correspondiente de Transaction
			switch i {
			case 0:
				newTransaction.Id = string(bytes.Trim(transactionBytes, "\x00"))
			case 1:
				mapKey = string(bytes.Trim(transactionBytes, "\x00")) // BuyerId
			case 2:
				currentIp := string(bytes.Trim(transactionBytes, "\x00"))
				foundOrigin, exists := originsMap[currentIp]
				if exists {
					newTransaction.Origin = foundOrigin
				} else {
					newOrigin := new(Origin)
					newOrigin.Ip = currentIp
					newOrigin.DType = "Origin"
					originsMap[currentIp] = newOrigin
					originsToDb = append(originsToDb, newOrigin)
					newTransaction.Origin = newOrigin
				}
			case 3:
				if newTransaction.Origin.Device == "" {
					newTransaction.Origin.Device = string(bytes.Trim(transactionBytes, "\x00"))
				}
			case 4:
				productIds := string(bytes.Trim(transactionBytes, "\x00"))
				productIds = productIds[1 : len(productIds)-1] // Eliminación de paréntesis
				productIdsArray := strings.Split(productIds, ",")
				var productsArray []*Product
				for _, product := range productIdsArray {
					productsArray = append(productsArray, productsMap[product])
				}
				newTransaction.Products = productsArray
			}

		}

		// Uso de las flags para control de for forever
		if loopFlag {
			break
		}
		if emptyFlag {
			continue
		}

		// Agrega la transacción al map de Transactions
		newTransaction.DType = "Transaction"
		_, exists := transactionsMap[newTransaction.Id] // Evitando duplicados...
		if !exists {
			transactionsMap[newTransaction.Id] = newTransaction
			transactionsDataProcessed[mapKey] = append(transactionsDataProcessed[mapKey], newTransaction)
		}

	}

	originsToDbChannel <- originsToDb
	close(originsToDbChannel)
	originsChannel <- originsMap
	close(originsChannel)
	transactionsChannel <- transactionsDataProcessed
	close(transactionsChannel)

}

func fetchBuyers(date int64, transactionsChannel <-chan map[string][]*Transaction, buyersChannel chan<- []*Buyer) {

	responseBuyers := utils.GenericFetch("buyers", date)

	defer responseBuyers.Body.Close()

	// Uso de ReadAll para leer todo el Body del http.Get (se necesita todo)
	buyersData, errBuyersData := io.ReadAll(responseBuyers.Body)
	utils.PanicErrorHandler(errBuyersData)

	var buyersDataProcessed []*Buyer
	unmarshallError := utils.JsoniterUnmarshall(buyersData, &buyersDataProcessed, "external")
	utils.PanicErrorHandler(unmarshallError)

	// Quitando duplicados de la data recibida y parseada...
	buyersMap := make(map[string]*Buyer)
	for _, buyer := range buyersDataProcessed {
		_, exists := buyersMap[buyer.Id]
		if !exists {
			buyersMap[buyer.Id] = buyer
		}
	}
	buyersDataProcessed = nil
	for _, buyer := range buyersMap {
		buyersDataProcessed = append(buyersDataProcessed, buyer)
	}

	transactionsMap := <-transactionsChannel
	// Cruzando transacciones con sus compradores:
	for _, buyer := range buyersDataProcessed {
		buyer.DType = "Buyer"
		foundTransactions, exist := transactionsMap[buyer.Id]
		if exist {
			buyer.Transactions = foundTransactions
		}
	}

	buyersChannel <- buyersDataProcessed
	close(buyersChannel)

}

func FetchDayData(date int64, queryProducts string, queryOrigins string) []*Buyer {

	productsChannel := make(chan map[string]*Product)
	productsToDbChannel := make(chan []*Product)
	productsFromDbChannel := make(chan *ProductQuery)
	transactionsChannel := make(chan map[string][]*Transaction)
	originsChannel := make(chan map[string]*Origin)
	originsToDbChannel := make(chan []*Origin)
	originsFromDbChannel := make(chan *OriginQuery)
	buyersChannel := make(chan []*Buyer)

	go fetchProducts(date, productsChannel, productsToDbChannel)
	go initialProductsOperation(queryProducts, productsToDbChannel, productsFromDbChannel)
	go fetchTransactions(date, productsChannel, transactionsChannel, originsChannel, originsToDbChannel)
	go initialOriginsOperation(queryOrigins, originsToDbChannel, originsFromDbChannel)
	go fetchBuyers(date, transactionsChannel, buyersChannel)

	var wg sync.WaitGroup

	wg.Add(1)
	go func(productsFromDbChannel chan *ProductQuery, productsChannel chan map[string]*Product) {

		defer wg.Done()
		productsWithUid := <-productsFromDbChannel
		productsMap := <-productsChannel

		for _, product := range productsWithUid.Q {
			foundProduct := productsMap[product.Id]
			foundProduct.Uid = product.Uid
		}

	}(productsFromDbChannel, productsChannel)

	wg.Add(1)
	go func(originsFromDbChannel chan *OriginQuery, originsChannel chan map[string]*Origin) {

		defer wg.Done()
		originsWithUid := <-originsFromDbChannel
		originsMap := <-originsChannel

		for _, origin := range originsWithUid.Q {
			foundOrigin := originsMap[origin.Ip]
			foundOrigin.Uid = origin.Uid
		}

	}(originsFromDbChannel, originsChannel)

	dayDataProcessed := <-buyersChannel

	return dayDataProcessed

}
