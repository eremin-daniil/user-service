package user

import (
	"github.com/google/uuid"
	"user_service/domain/entity/user/primitive"
	profileEntity "user_service/domain/entity/user/profile"
)

type ID struct {
	uuid.UUID
}

type User struct {
	id      *ID
	login   primitive.Login
	profile *profileEntity.Profile
}

func (u *User) Id() *ID {
	return u.id
}

func (u *User) Login() primitive.Login {
	return u.login
}

func (u *User) Profile() *profileEntity.Profile {
	return u.profile
}

func (u *User) SetProfile(profile *profileEntity.Profile) {
	u.profile = profile
}
