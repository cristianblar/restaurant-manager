package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"strconv"
	"strings"
)

func fetchProducts() *ProductsData {

	responseProducts := genericFetch("products")

	// Creación de reader con csv para leer csv
	csvReader := csv.NewReader(responseProducts.Body)
	// Configuración de csvReader para indicar separador
	csvReader.Comma = '\''
	// Uso de ReadAll para leer todo el csv (se necesita todo)
	productsData, errProductsData := csvReader.ReadAll()
	errorHandler(errProductsData)

	// Instanciación de struct ProductsData
	productsDataProcessed := new(ProductsData)
	productsDataProcessed.Products = make(map[string]*Product)
	for _, product := range productsData {
		// Conversión del precio del producto a int
		numPrice, errnumPrice := strconv.Atoi(product[2])
		errorHandler(errnumPrice)
		// Conversión del precio del producto en centavos a dólares
		var dollarPrice float32 = float32(numPrice) / 100
		newProduct := new(Product)
		newProduct.Id = product[0]
		newProduct.Name = product[1]
		newProduct.Price = dollarPrice
		productsDataProcessed.Products[newProduct.Id] = newProduct
	}

	return productsDataProcessed

}

func fetchTransactions(products map[string]*Product) *TransactionsData {
	responseTransactions := genericFetch("transactions")

	// Creación de reader con bufio para leer por bytes
	transactionsReader := bufio.NewReader(responseTransactions.Body)

	// Instanciación de struct TransactionsData
	transactionsDataProcessed := new(TransactionsData)
	transactionsDataProcessed.Transactions = make(map[string][]*Transaction)

	// for forever para la cantidad de transacciones que es desconocida
	for {
		// flags para detener o continuar for forever desde for interno
		loopFlag := false
		emptyFlag := false

		// Instanciación de struct Transaction para agregar a TransactionsData.Transactions
		newTransaction := new(Transaction)
		var mapKey string = ""

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
					log.Fatal(errTransactionBytes.Error())
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
				newTransaction.Ip = string(bytes.Trim(transactionBytes, "\x00"))
			case 3:
				newTransaction.Device = string(bytes.Trim(transactionBytes, "\x00"))
			case 4:
				productIds := string(bytes.Trim(transactionBytes, "\x00"))
				productIds = productIds[1 : len(productIds)-1] // Eliminación de paréntesis
				productIdsArray := strings.Split(productIds, ",")
				var productsArray []*Product
				for _, product := range productIdsArray {
					productsArray = append(productsArray, products[product])
				}
				newTransaction.Products = productsArray
			}

		}

		// Uso de las flags para manejar el for forever
		if loopFlag {
			break
		}
		if emptyFlag {
			continue
		}

		// Agrega la transacción al map de TransactionsData.Transactions
		transactionsDataProcessed.Transactions[mapKey] = append(transactionsDataProcessed.Transactions[mapKey], newTransaction)

	}

	return transactionsDataProcessed
}

func fetchBuyers(transactions map[string][]*Transaction) *BuyersData {

	responseBuyers := genericFetch("buyers")

	// Uso de ReadAll para leer todo el Body del http.Get (se necesita todo)
	buyersData, errBuyersData := io.ReadAll(responseBuyers.Body)
	errorHandler(errBuyersData)

	// JSON Unmarshal (paso de JSON en string a Go structs)
	// Instanciación de struct BuyersData
	buyersDataProcessed := new(BuyersData)
	json.Unmarshal(buyersData, &buyersDataProcessed.Buyers)

	for _, buyer := range buyersDataProcessed.Buyers {
		foundTransactions, exist := transactions[buyer.Id]
		if exist {
			buyer.Transactions = foundTransactions
		}
	}

	return buyersDataProcessed
}
