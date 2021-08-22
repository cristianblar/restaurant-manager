package main

type Product struct {
	Id    string  `dgraph:"Product.id,omitempty"`
	Name  string  `dgraph:"Product.name,omitempty"`
	Price float32 `dgraph:"Product.price,omitempty"`
	DType string  `dgraph:"dgraph.type,omitempty"`
}

type Origin struct {
	Ip     string `dgraph:"Origin.ip,omitempty"`
	Device string `dgraph:"Origin.device,omitempty"`
	DType  string `dgraph:"dgraph.type,omitempty"`
}

type Transaction struct {
	Id       string     `dgraph:"Transaction.id,omitempty"`
	Origin   *Origin    `dgraph:"Transaction.origin,omitempty"`
	Products []*Product `dgraph:"Transaction.products,omitempty"`
	DType    string     `dgraph:"dgraph.type,omitempty"`
}

type Buyer struct {
	Id           string         `external:"id,omitempty" dgraph:"Buyer.id,omitempty"`
	Name         string         `external:"name,omitempty" dgraph:"Buyer.name,omitempty"`
	Age          uint8          `external:"age,omitempty" dgraph:"Buyer.age,omitempty"`
	Transactions []*Transaction `dgraph:"Buyer.transactions,omitempty"`
	DType        string         `dgraph:"dgraph.type,omitempty"`
}
