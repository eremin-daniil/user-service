package usecase

import (
	"context"
	"user_service/boundary/dto"
)

type ProfileUseCaseInterface interface {
	Create(ctx context.Context, profile *dto.ProfileDTO) (string, error)
	GetByID(ctx context.Context, id string) (*dto.ProfileDTO, error)
	Update(ctx context.Context, profile *dto.ProfileDTO) error
}
