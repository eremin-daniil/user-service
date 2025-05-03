package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"user_service/adapter/repository/model"
	"user_service/common/repository"
	userEntity "user_service/domain/entity/user"
	userPrimitive "user_service/domain/entity/user/primitive"
	mongoInterface "user_service/infrastructure/mongo/interface"
)

type UserRepository struct {
	collection string
	mongo      mongoInterface.MongoRepositoryInterface
}

func NewUserRepository(mongo mongoInterface.MongoRepositoryInterface) *UserRepository {
	return &UserRepository{
		collection: "user",
		mongo:      mongo,
	}
}

func (u *UserRepository) Create(ctx context.Context, user *userEntity.User) (repository.ObjectID, error) {
	objectID, err := u.mongo.Insert(ctx, u.collection, model.UserModelFromEntity(user))
	if err != nil {
		return "", err
	}
	return repository.ObjectID(objectID), nil
}

func (u *UserRepository) GetByID(ctx context.Context, userID userEntity.ID) (*userEntity.User, error) {
	query := bson.D{{"user_id", userID.String()}}
	user := model.NewDefaultUserModel()
	ok, err := u.mongo.FindOne(ctx, u.collection, user, query)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	entity, err := user.ToEntity()
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (u *UserRepository) GetByLogin(ctx context.Context, login userPrimitive.Login) (*userEntity.User, error) {
	query := bson.D{{"login", login.String()}}
	user := model.NewDefaultUserModel()
	ok, err := u.mongo.FindOne(ctx, u.collection, user, query)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	entity, err := user.ToEntity()
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (u *UserRepository) Update(ctx context.Context, user *userEntity.User) error {
	filter := bson.D{{"user_id", user.Id().String()}}
	userModel := model.UserModelFromEntity(user)
	ok, err := u.mongo.ReplaceOne(ctx, u.collection, filter, userModel)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	return nil
}

func (u *UserRepository) DeleteByID(ctx context.Context, id userEntity.ID) error {
	filter := bson.D{{"user_id", id.String()}}
	ok, err := u.mongo.DeleteOne(ctx, u.collection, filter)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	return nil
}
