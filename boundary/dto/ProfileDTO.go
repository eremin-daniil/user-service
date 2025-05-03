package dto

import (
	"time"
	profileEntity "user_service/domain/entity/user/profile"
)

type ProfileDTO struct {
	ID              string
	LastName        string
	FirstName       string
	CreatedDateTime time.Time
	UpdatedDateTime time.Time
}

func ProfileToDTO(entity *profileEntity.Profile) *ProfileDTO {
	return &ProfileDTO{
		ID:              entity.Id().String(),
		LastName:        entity.LastName(),
		FirstName:       entity.FirstName(),
		CreatedDateTime: entity.CreatedDateTime(),
		UpdatedDateTime: entity.UpdatedDateTime(),
	}
}

func (dto *ProfileDTO) ToEntity() (*profileEntity.Profile, error) {
	return profileEntity.NewBuilder().
		IDString(dto.ID).
		LastName(dto.LastName).
		FirstName(dto.FirstName).
		CreatedDateTime(dto.CreatedDateTime).
		UpdatedDateTime(dto.UpdatedDateTime).
		Build()
}
