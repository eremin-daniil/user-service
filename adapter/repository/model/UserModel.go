package model

import (
	userEntity "user_service/domain/entity/user"
	profileEntity "user_service/domain/entity/user/profile"
)

type UserModel struct {
	ID      string        `bson:"user_id"`
	Login   string        `bson:"login"`
	Profile *ProfileModel `bson:"profile"`
}

func NewDefaultUserModel() *UserModel {
	return &UserModel{
		ID:      "",
		Login:   "",
		Profile: nil,
	}
}

func UserModelFromEntity(entity *userEntity.User) *UserModel {
	var profile *ProfileModel
	if entity.Profile() != nil {
		profile = ProfileModelFromEntity(entity.Profile())
	}
	return &UserModel{
		ID:      entity.Id().String(),
		Login:   entity.Login().String(),
		Profile: profile,
	}
}

func (m *UserModel) ToEntity() (*userEntity.User, error) {
	profile, err := m.getProfile()
	if err != nil {
		return nil, err
	}
	return userEntity.NewBuilder().
		IDString(m.ID).
		LoginString(m.Login).
		Profile(profile).
		Build()
}

func (m *UserModel) getProfile() (*profileEntity.Profile, error) {
	if m.Profile == nil {
		return nil, nil
	}
	return m.Profile.ToEntity()
}
