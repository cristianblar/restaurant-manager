package main

const currentDomain string = "http://localhost:3000"

const httpErrorMessage = "Something went wrong... Please, try again"

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
	Product.id: string @index(hash) .
	Product.name: string .
	Product.price: float .
	Origin.ip: string @index(hash) .
	Origin.device: string .
	Transaction.id: string @index(hash) .
	Transaction.origin: uid @reverse .
	Transaction.products: [uid] @count .
	Buyer.id: string @index(hash) .
	Buyer.name: string .
	Buyer.age: int .
	Buyer.transactions: [uid] @count @reverse .
`

const queryProducts string = `
	{
		q(func: type(Product)){
			uid
			Product.id
			Product.name
			Product.price
		}
	}
`

const queryOrigins string = `
	{
		q(func: type(Origin)){
			uid
			Origin.ip
			Origin.device
		}
	}
`

const queryAllBuyers string = `
	query q(){
		var(func: type(Buyer)) @filter(has(Buyer.transactions)){
			numTransactions as count(Buyer.transactions)
		}
		q(func: type(Buyer), orderdesc: val(numTransactions)) @filter(has(Buyer.transactions)){
			Buyer.id
			Buyer.name
			Buyer.age
			Buyer.transactionsAmount: val(numTransactions)
		}
	}
`

const queryBuyerById string = `
	query q($id: string){
		owner(func: eq(Buyer.id, $id)){
			Buyer.id
			OwnerName AS Buyer.name
			Buyer.age
			Buyer.transactions (orderdesc: Transaction.id){
				Transaction.id
				OwnerIPs AS Transaction.origin{
					Origin.ip
					Origin.device
				}
				OwnerProducts AS Transaction.products (orderasc: Product.price){
					Product.id
					Product.name
					Product.price
				}
			}
		}
		otherBuyers(func: uid(OwnerIPs)) @normalize{
			CommonTransactions AS ~Transaction.origin{
				~Buyer.transactions @filter(NOT eq(Buyer.name, val(OwnerName))){
					Buyer.id: Buyer.id
					Buyer.name: Buyer.name
					Buyer.age: Buyer.age
				}
			}
		}
		otherProducts(func: uid(CommonTransactions))@normalize{
			Transaction.products (orderdesc: Product.price , first: 1) @filter(NOT uid(OwnerProducts)){
				Product.id: Product.id
				Product.name: Product.name
				Product.price: Product.price
			}
		}
	}
`
