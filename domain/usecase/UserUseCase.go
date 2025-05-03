package usecase

import (
	"context"
	"github.com/google/uuid"
	"user_service/boundary/dto"
	"user_service/boundary/gateway"
	"user_service/boundary/repository"
	"user_service/domain/entity/user"
	"user_service/domain/entity/user/primitive"
)

type UserUseCase struct {
	repository repository.UserRepositoryInterface
	gateway    gateway.UserEventDrivenGatewayInterface
}

func (u *UserUseCase) Register(ctx context.Context, user *dto.UserDTO) (string, error) {
	entity, err := user.ToEntity()
	if err != nil {
		return "", err
	}
	_, err = u.repository.Create(ctx, entity)
	if err != nil {
		return "", err
	}
	err = u.gateway.SendCreateUserEvent(ctx, dto.UserToDTO(entity))
	if err != nil {
		return "", err
	}
	return entity.Id().String(), nil
}

func (u *UserUseCase) GetByID(ctx context.Context, id string) (*dto.UserDTO, error) {
	entityID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	entity, err := u.repository.GetByID(ctx, user.ID{UUID: entityID})
	if err != nil {
		return nil, err
	}
	return dto.UserToDTO(entity), nil
}

func (u *UserUseCase) GetByLogin(ctx context.Context, login string) (*dto.UserDTO, error) {
	loginPrimitive, err := primitive.LoginFromString(login)
	if err != nil {
		return nil, err
	}
	entity, err := u.repository.GetByLogin(ctx, loginPrimitive)
	if err != nil {
		return nil, err
	}
	return dto.UserToDTO(entity), nil
}
