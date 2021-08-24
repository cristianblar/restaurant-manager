package main

import (
	"context"

	dgo "github.com/dgraph-io/dgo/v210"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

// Database

type DatabaseConnection struct {
	Connection   *grpc.ClientConn
	DgraphClient *dgo.Dgraph
	Context      context.Context
}

// Web server

type Pagination struct {
	TotalResults int     `dgraph:"Buyers.total,omitempty"`
	TotalPages   int     `dgraph:"Pages.total,omitempty"`
	PreviousPage string  `dgraph:"Page.previous,omitempty"`
	NextPage     string  `dgraph:"Page.next,omitempty"`
	Results      []Buyer `dgraph:"results,omitempty"`
}

type Server struct {
	Port   string
	Router *chi.Mux
}

// Queries

type ProductQuery struct {
	Q []Product `dgraph:"q,omitempty"`
}

type OriginQuery struct {
	Q []Origin `dgraph:"q,omitempty"`
}

type AllBuyersQuery struct {
	Q []Buyer `dgraph:"q,omitempty"`
}

type BuyerQuery struct {
	Owner         []Buyer   `dgraph:"owner,omitempty"`
	OtherBuyers   []Buyer   `dgraph:"otherBuyers,omitempty"`
	OtherProducts []Product `dgraph:"otherProducts,omitempty"`
}

// Data

type Product struct {
	Uid   string  `dgraph:"uid,omitempty"`
	Id    string  `dgraph:"Product.id,omitempty"`
	Name  string  `dgraph:"Product.name,omitempty"`
	Price float32 `dgraph:"Product.price,omitempty"`
	DType string  `dgraph:"dgraph.type,omitempty"`
}

type Origin struct {
	Uid    string `dgraph:"uid,omitempty"`
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
	Id                 string         `external:"id,omitempty" dgraph:"Buyer.id,omitempty"`
	Name               string         `external:"name,omitempty" dgraph:"Buyer.name,omitempty"`
	Age                uint8          `external:"age,omitempty" dgraph:"Buyer.age,omitempty"`
	Transactions       []*Transaction `dgraph:"Buyer.transactions,omitempty"`
	TransactionsAmount uint8          `dgraph:"Buyer.transactionsAmount,omitempty"`
	DType              string         `dgraph:"dgraph.type,omitempty"`
}
