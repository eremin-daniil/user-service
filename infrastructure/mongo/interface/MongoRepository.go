package _interface

import (
	"context"
)

type MongoRepositoryInterface interface {
	Insert(ctx context.Context, collection string, data interface{}) (string, error)
	ReplaceOne(ctx context.Context, collection string, filter, data interface{}) (bool, error)
	FindOne(ctx context.Context, collection string, result, query interface{}) (bool, error)
	DeleteOne(ctx context.Context, collection string, filter interface{}) (bool, error)
}
