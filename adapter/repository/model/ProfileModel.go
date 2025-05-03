package model

import (
	"time"
	profileEntity "user_service/domain/entity/user/profile"
)

type ProfileModel struct {
	ID              string    `bson:"profile_id"`
	LastName        string    `bson:"last_name"`
	FirstName       string    `bson:"first_name"`
	CreatedDateTime time.Time `bson:"created_date_time"`
	UpdatedDateTime time.Time `bson:"updated_date_time"`
}

func ProfileModelFromEntity(entity *profileEntity.Profile) *ProfileModel {
	return &ProfileModel{
		ID:              entity.Id().String(),
		LastName:        entity.LastName(),
		FirstName:       entity.FirstName(),
		CreatedDateTime: entity.CreatedDateTime(),
		UpdatedDateTime: entity.UpdatedDateTime(),
	}
}

func (m *ProfileModel) ToEntity() (*profileEntity.Profile, error) {
	return profileEntity.NewBuilder().
		IDString(m.ID).
		LastName(m.LastName).
		FirstName(m.FirstName).
		CreatedDateTime(m.CreatedDateTime).
		UpdatedDateTime(m.UpdatedDateTime).
		Build()
}
