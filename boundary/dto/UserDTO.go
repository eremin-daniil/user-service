package dto

import (
	userEntity "user_service/domain/entity/user"
	profileEntity "user_service/domain/entity/user/profile"
)

type UserDTO struct {
	ID      string
	Login   string
	Profile *ProfileDTO
}

func UserToDTO(entity *userEntity.User) *UserDTO {
	var profile *ProfileDTO
	if entity.Profile() != nil {
		profile = ProfileToDTO(entity.Profile())
	}
	return &UserDTO{
		ID:      entity.Id().String(),
		Login:   entity.Login().String(),
		Profile: profile,
	}
}

func (d *UserDTO) ToEntity() (*userEntity.User, error) {
	profile, err := d.getProfile()
	if err != nil {
		return nil, err
	}
	return userEntity.NewBuilder().
		IDString(d.ID).
		LoginString(d.Login).
		Profile(profile).
		Build()
}

func (d *UserDTO) getProfile() (*profileEntity.Profile, error) {
	if d.Profile != nil {
		return d.Profile.ToEntity()
	}
	return nil, nil
}
