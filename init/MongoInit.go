package init

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func MongoDatabase(url, mongoDBName string) (*mongo.Database, error) {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	pingTimeout := time.Now().Add(1 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), pingTimeout)
	defer cancelFunc()
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return mongoClient.Database(mongoDBName), nil
}
