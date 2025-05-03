package usecase

import (
	"context"
	"github.com/google/uuid"
	"user_service/boundary/dto"
	"user_service/boundary/gateway"
	"user_service/boundary/repository"
	profileEntity "user_service/domain/entity/user/profile"
)

type ProfileUseCase struct {
	repository repository.ProfileRepositoryInterface
	gateway    gateway.ProfileEventDrivenGatewayInterface
}

func (p *ProfileUseCase) Create(ctx context.Context, profile *dto.ProfileDTO) (string, error) {
	entity, err := profile.ToEntity()
	if err != nil {
		return "", err
	}
	_, err = p.repository.Create(ctx, entity)
	if err != nil {
		return "", err
	}
	err = p.gateway.SendCreateProfileEvent(ctx, dto.ProfileToDTO(entity))
	if err != nil {
		return "", err
	}
	return entity.Id().String(), nil
}

func (p *ProfileUseCase) GetByID(ctx context.Context, id string) (*dto.ProfileDTO, error) {
	entityID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	entity, err := p.repository.GetByID(ctx, profileEntity.ID{UUID: entityID})
	if err != nil {
		return nil, err
	}
	return dto.ProfileToDTO(entity), nil
}

func (p *ProfileUseCase) Update(ctx context.Context, profile *dto.ProfileDTO) error {
	entity, err := profile.ToEntity()
	if err != nil {
		return err
	}
	err = p.repository.Update(ctx, entity)
	if err != nil {
		return err
	}
	err = p.gateway.SendUpdateProfileEvent(ctx, dto.ProfileToDTO(entity))
	if err != nil {
		return err
	}
	return nil
}
