package main

type Product struct {
	Id    string
	Name  string
	Price float32
}

type ProductsData struct {
	Products map[string]*Product
}

type Transaction struct {
	Id       string
	Ip       string
	Device   string
	Products []*Product
}

type TransactionsData struct {
	Transactions map[string][]*Transaction
}

type Buyer struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Age          uint8  `json:"age"`
	Transactions []*Transaction
}

type BuyersData struct {
	Buyers []*Buyer
}
