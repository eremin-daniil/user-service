package usecase

import (
	"context"
	"user_service/boundary/dto"
)

type UserUseCaseInterface interface {
	RegisterUser(ctx context.Context, user *dto.UserDTO) (string, error)
	GetUserByID(ctx context.Context, id string) (*dto.UserDTO, error)
	GetUserByLogin(ctx context.Context, login string) (*dto.UserDTO, error)

	CreateProfile(ctx context.Context, userID string, profile *dto.ProfileDTO) (string, error)
	GetProfileByUserID(ctx context.Context, userID string) (*dto.ProfileDTO, error)
	UpdateProfile(ctx context.Context, userID string, profile *dto.ProfileDTO) error
}
