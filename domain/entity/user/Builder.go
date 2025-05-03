package user

import (
	"errors"
	"github.com/google/uuid"
	"user_service/domain/entity/user/primitive"
	profileEntity "user_service/domain/entity/user/profile"
)

var (
	ErrInvalidUserID   = errors.New("invalid user ID")
	ErrLoginIsRequired = errors.New("login is required")
)

type Builder struct {
	id      *ID
	login   primitive.Login
	profile *profileEntity.Profile

	errs []error
}

func NewBuilder() *Builder {
	return &Builder{
		id:      nil,
		login:   "",
		profile: nil,
		errs:    make([]error, 0),
	}
}

func (b *Builder) IDString(value string) *Builder {
	if value == "" {
		return b
	}
	id, err := uuid.Parse(value)
	if err != nil {
		b.errs = append(b.errs, ErrInvalidUserID)
		return b
	}
	b.id = &ID{id}
	return b
}

func (b *Builder) LoginString(value string) *Builder {
	login, err := primitive.LoginFromString(value)
	if err != nil {
		b.errs = append(b.errs, err)
		return b
	}
	b.login = login
	return b
}

func (b *Builder) Profile(profile *profileEntity.Profile) *Builder {
	b.profile = profile
	return b
}

func (b *Builder) Build() (*User, error) {
	b.fillDefaultFields()
	b.checkRequiredFields()
	if len(b.errs) > 0 {
		return nil, errors.Join(b.errs...)
	}
	user := &User{
		id:      b.id,
		login:   b.login,
		profile: b.profile,
	}
	return user, nil
}

func (b *Builder) fillDefaultFields() {
	if b.id == nil {
		b.id = &ID{uuid.New()}
	}
}

func (b *Builder) checkRequiredFields() {
	if b.login == "" {
		b.errs = append(b.errs, ErrLoginIsRequired)
	}
}
