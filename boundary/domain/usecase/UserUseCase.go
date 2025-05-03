package usecase

import (
	"context"
	"user_service/boundary/dto"
)

type UserUseCaseInterface interface {
	Register(ctx context.Context, user *dto.UserDTO) (string, error)
	GetByID(ctx context.Context, id string) (*dto.UserDTO, error)
	GetByLogin(ctx context.Context, login string) (*dto.UserDTO, error)
}
