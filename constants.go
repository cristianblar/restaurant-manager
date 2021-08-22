package main

const schemaObject string = `
	type Product {
		Product.id
		Product.name
		Product.price
	}
	type Origin {
		Origin.ip
		Origin.device
	}
	type Transaction {
		Transaction.id
		Transaction.origin
		Transaction.products
	}
	type Buyer {
		Buyer.id
		Buyer.name
		Buyer.age
		Buyer.transactions
	}
	Product.id: string @index(hash) @upsert .
	Product.name: string .
	Product.price: float .
	Origin.ip: string @index(hash) @upsert .
	Origin.device: string .
	Transaction.id: string @index(hash) @upsert .
	Transaction.origin: uid .
	Transaction.products: [uid] @count .
	Buyer.id: string @index(hash) @upsert .
	Buyer.name: string .
	Buyer.age: int .
	Buyer.transactions: [uid] @count .
`

const queryAllBuyers string = `
	{
		q(func: type(Buyer)) @filter(has(Buyer.transactions)){
			Buyer.id
			Buyer.name
			Buyer.age
		}
	}
`
