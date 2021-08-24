package database

import (
	"context"

	"github.com/dgraph-io/dgo/v210"
	"google.golang.org/grpc"
)

type DatabaseConnection struct {
	Connection   *grpc.ClientConn
	DgraphClient *dgo.Dgraph
	Context      context.Context
}
