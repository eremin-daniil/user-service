package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	loggerInterface "user_service/infrastructure/logger/interface"
)

type Repository struct {
	mongoDb *mongo.Database
	logger  loggerInterface.Logger
}

func NewMongoRepository(mongoDb *mongo.Database, logger loggerInterface.Logger) *Repository {
	return &Repository{mongoDb: mongoDb, logger: logger}
}

func (r *Repository) Insert(ctx context.Context, collection string, data interface{}) (string, error) {
	collectionResult := r.mongoDb.Collection(collection)
	result, err := collectionResult.InsertOne(ctx, data)
	if err != nil {
		r.logger.Error(ctx, err)
		return "", err
	}
	var id string
	if result != nil {
		resID, ok := result.InsertedID.(primitive.ObjectID)
		if ok {
			id = resID.Hex()
		}
	}
	return id, nil
}

func (r *Repository) ReplaceOne(ctx context.Context, collection string, filter interface{}, data interface{}) (bool, error) {
	collectionResult := r.mongoDb.Collection(collection)
	result, err := collectionResult.ReplaceOne(ctx, filter, data, options.Replace())
	if err != nil {
		r.logger.Error(ctx, err)
		return false, err
	}
	return result.ModifiedCount > 0, nil
}

func (r *Repository) FindOne(ctx context.Context, collection string, resultModel, findQuery interface{}) (bool, error) {
	collectionResult := r.mongoDb.Collection(collection)
	result := collectionResult.FindOne(ctx, findQuery, options.FindOne())
	if result.Err() != nil {
		r.logger.Error(ctx, result.Err())
		return false, result.Err()
	}
	err := result.Decode(resultModel)
	if err != nil {
		r.logger.Error(ctx, err)
		return false, err
	}
	return true, nil
}

func (r *Repository) DeleteOne(ctx context.Context, collection string, filter interface{}) (bool, error) {
	collectionResult := r.mongoDb.Collection(collection)
	result, err := collectionResult.DeleteOne(ctx, filter, options.Delete())
	if err != nil {
		r.logger.Error(ctx, err)
		return false, err
	}
	return result.DeletedCount > 0, nil
}
