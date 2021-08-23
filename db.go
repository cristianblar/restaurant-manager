package main

import (
	"context"
	"os"

	dgo "github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

type DatabaseConnection struct {
	Connection   *grpc.ClientConn
	DgraphClient *dgo.Dgraph
	Context      context.Context
}

// DB Singleton:
var databaseSingleton *DatabaseConnection = nil

func CreateDatabase(schema string) *DatabaseConnection {
	if databaseSingleton == nil {
		databaseSingleton = new(DatabaseConnection)
		ctx := context.Background()
		databaseSingleton.Context = ctx
		databaseSingleton.connectToDb()
		databaseSingleton.applySchema(schema)
		databaseSingleton.Connection.Close()
		databaseSingleton.Connection = nil
		databaseSingleton.DgraphClient = nil
	}

	return databaseSingleton

}

func (db *DatabaseConnection) connectToDb() {

	if db.Connection == nil {
		DGRAPH_ENDPOINT := os.Getenv("DGRAPH_ENDPOINT")
		dialOpts := append([]grpc.DialOption{}, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
		conn, connError := grpc.Dial(DGRAPH_ENDPOINT, dialOpts...)
		errorHandler(connError)

		dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

		db.Connection = conn
		db.DgraphClient = dgraphClient
	}

}

func (db *DatabaseConnection) applySchema(schema string) {

	// Limpiar la DB por primera vez:
	operationObject := &api.Operation{
		DropOp: api.Operation_ALL,
	}
	// Ejecución de operación:
	operationError := db.DgraphClient.Alter(db.Context, operationObject)
	errorHandler(operationError)
	// Schema object:
	schemaObject := &api.Operation{
		Schema: schema,
	}
	// Ejecución de operación:
	schemaError := db.DgraphClient.Alter(db.Context, schemaObject)
	errorHandler(schemaError)

}

func (db *DatabaseConnection) DropData() {

	// Operation object:
	operationObject := &api.Operation{
		DropOp: api.Operation_DATA,
	}
	// Ejecución de operación:
	operationError := db.DgraphClient.Alter(db.Context, operationObject)
	errorHandler(operationError)

}

func (db *DatabaseConnection) BulkJsonMutation(transactionsList []byte) {

	db.connectToDb()

	txn := db.DgraphClient.NewTxn()
	// Mutation object:
	mutationObject := &api.Mutation{
		SetJson: transactionsList,
	}
	// Ejecución de mutation:
	_, mutationError := txn.Mutate(db.Context, mutationObject)
	if mutationError != nil {
		txn.Discard(db.Context)
		errorHandler(mutationError)
	}
	// Commit de la transacción:
	commitError := txn.Commit(db.Context)
	if commitError != nil {
		txn.Discard(db.Context)
		errorHandler(commitError)
	}

	db.Connection.Close()
	db.Connection = nil
	db.DgraphClient = nil

}

func (db *DatabaseConnection) GetQuery(query string) []byte {

	db.connectToDb()

	txn := db.DgraphClient.NewTxn()

	queryResponse, queryError := txn.Query(db.Context, query)
	errorHandler(queryError)

	queryJson := queryResponse.GetJson()

	db.Connection.Close()
	db.Connection = nil
	db.DgraphClient = nil

	return queryJson

}

func (db *DatabaseConnection) GetQueryWithVariables(query string, vars map[string]string) []byte {

	db.connectToDb()

	txn := db.DgraphClient.NewTxn()

	queryResponse, queryError := txn.QueryWithVars(db.Context, query, vars)
	errorHandler(queryError)

	queryJson := queryResponse.GetJson()

	db.Connection.Close()
	db.Connection = nil
	db.DgraphClient = nil

	return queryJson

}
