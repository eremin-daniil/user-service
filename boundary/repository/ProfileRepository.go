package repository

import (
	"context"
	"user_service/common/repository"
	profileEntity "user_service/domain/entity/user/profile"
)

type ProfileRepositoryInterface interface {
	Create(ctx context.Context, profile *profileEntity.Profile) (repository.ObjectID, error)
	GetByID(ctx context.Context, id profileEntity.ID) (*profileEntity.Profile, error)
	Update(ctx context.Context, profile *profileEntity.Profile) error
	DeleteByID(ctx context.Context, id profileEntity.ID) error
}
