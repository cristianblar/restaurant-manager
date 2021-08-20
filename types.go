package main

type Product struct {
	Id    string
	Name  string
	Price float32
}

type Transaction struct {
	Id       string
	Ip       string
	Device   string
	Products []*Product
}

type Buyer struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Age          uint8  `json:"age"`
	Transactions []*Transaction
}

type Day struct {
	Date   int64
	Buyers []*Buyer
}
