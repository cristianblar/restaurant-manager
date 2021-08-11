package main

type Product struct {
	Id    string
	Name  string
	Price float32
}

type ProductsData struct {
	Products []*Product
}

type Buyer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type BuyersData struct {
	Buyers []*Buyer
}

type Transaction struct {
	Id         string
	BuyerId    string
	Ip         string
	Device     string
	ProductIds []string
}

type TransactionsData struct {
	Transactions []*Transaction
}
