package repository

import (
	"context"
	"user_service/common/repository"
	userEntity "user_service/domain/entity/user"
	"user_service/domain/entity/user/primitive"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *userEntity.User) (repository.ObjectID, error)
	GetByID(ctx context.Context, userID userEntity.ID) (*userEntity.User, error)
	GetByLogin(ctx context.Context, login primitive.Login) (*userEntity.User, error)
	Update(ctx context.Context, user *userEntity.User) error
	DeleteByID(ctx context.Context, userID userEntity.ID) error
}
