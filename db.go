package main

import (
	"context"
	"os"

	dgo "github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func connectToDb(firstConnection bool, newDate bool, dbOp func(dgraphClient *dgo.Dgraph, ctx context.Context, opData []byte) []byte, opData []byte) []byte {

	DGRAPH_ENDPOINT := os.Getenv("DGRAPH_ENDPOINT")
	dialOpts := append([]grpc.DialOption{}, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name), grpc.MaxCallRecvMsgSize((1024*6)*1024)))
	conn, connError := grpc.Dial(DGRAPH_ENDPOINT, dialOpts...)
	errorHandler(connError)
	defer conn.Close()

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	ctx := context.Background()

	if firstConnection {
		return applySchema(dgraphClient, ctx) // Sincroniza esquema por primera vez (BORRA TODOS LOS DATOS)
	} else {
		if newDate {
			dropData(dgraphClient, ctx) // Nueva fecha sincronizada, elimina datos anteriores
		}
		return dbOp(dgraphClient, ctx, opData) // Mutation o Query
	}

}

func applySchema(dgraphClient *dgo.Dgraph, ctx context.Context) []byte {
	// Operation object:
	operationObject := &api.Operation{
		DropOp: api.Operation_ALL,
	}
	// Ejecución de operación:
	operationError := dgraphClient.Alter(ctx, operationObject)
	errorHandler(operationError)
	// Schema object:
	schemaObject := &api.Operation{
		Schema: schemaObject,
	}
	// Ejecución de operación:
	schemaError := dgraphClient.Alter(ctx, schemaObject)
	errorHandler(schemaError)

	return []byte("done")
}

func dropData(dgraphClient *dgo.Dgraph, ctx context.Context) []byte {
	// Operation object:
	operationObject := &api.Operation{
		DropOp: api.Operation_DATA,
	}
	// Ejecución de operación:
	operationError := dgraphClient.Alter(ctx, operationObject)
	errorHandler(operationError)

	return []byte("done")
}

func applyMutation(dgraphClient *dgo.Dgraph, ctx context.Context, opData []byte) []byte {
	txn := dgraphClient.NewTxn()
	// Mutation object:
	mutationObject := &api.Mutation{
		SetJson: opData,
	}
	// Ejecución de mutation:
	_, mutationError := txn.Mutate(ctx, mutationObject)
	if mutationError != nil {
		txn.Discard(ctx)
		errorHandler(mutationError)
	}
	// Commit de la transacción:
	commitError := txn.Commit(ctx)
	if commitError != nil {
		txn.Discard(ctx)
		errorHandler(commitError)
	}

	return []byte("done")
}

func getQuery(dgraphClient *dgo.Dgraph, ctx context.Context, opData []byte) []byte {
	txn := dgraphClient.NewTxn()

	queryResponse, queryError := txn.Query(ctx, string(opData))
	errorHandler(queryError)

	queryJson := queryResponse.GetJson()

	return queryJson

}
