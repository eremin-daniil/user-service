package profile

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidProfileID                 = errors.New("invalid profile ID")
	ErrLastNameIsRequired               = errors.New("last name is required")
	ErrFirstNameIsRequired              = errors.New("first name is required")
	ErrCreatedDateTimeProfileIsRequired = errors.New("created date time profile is required")
)

type Builder struct {
	id              *ID
	lastName        string
	firstName       string
	createdDateTime time.Time
	updatedDateTime time.Time

	errs []error
}

func NewBuilder() *Builder {
	return &Builder{
		id:              nil,
		lastName:        "",
		firstName:       "",
		createdDateTime: time.Time{},
		updatedDateTime: time.Time{},
		errs:            make([]error, 0),
	}
}

func (b *Builder) IDString(value string) *Builder {
	if value == "" {
		return b
	}
	id, err := uuid.Parse(value)
	if err != nil {
		b.errs = append(b.errs, ErrInvalidProfileID)
		return b
	}
	b.id = &ID{id}
	return b
}

func (b *Builder) LastName(lastName string) *Builder {
	if len(lastName) < 2 || len(lastName) > 32 {
		b.errs = append(b.errs, ErrIncorrectLengthLastName)
		return b
	}
	b.lastName = lastName
	return b
}

func (b *Builder) FirstName(firstName string) *Builder {
	if len(firstName) < 2 || len(firstName) > 32 {
		b.errs = append(b.errs, ErrIncorrectLengthFirstName)
		return b
	}
	b.firstName = firstName
	return b
}

func (b *Builder) CreatedDateTime(createdDateTime time.Time) *Builder {
	b.createdDateTime = createdDateTime
	return b
}

func (b *Builder) UpdatedDateTime(updatedDateTime time.Time) *Builder {
	b.updatedDateTime = updatedDateTime
	return b
}

func (b *Builder) Build() (*Profile, error) {
	b.fillDefaultFields()
	b.checkRequiredFields()
	if len(b.errs) > 0 {
		return nil, errors.Join(b.errs...)
	}
	profile := &Profile{
		id:              b.id,
		lastName:        b.lastName,
		firstName:       b.firstName,
		createdDateTime: b.createdDateTime,
		updatedDateTime: b.updatedDateTime,
	}
	return profile, nil
}

func (b *Builder) fillDefaultFields() {
	if b.id == nil {
		b.id = &ID{uuid.New()}
		b.createdDateTime = time.Now().UTC()
	}
}

func (b *Builder) checkRequiredFields() {
	if b.lastName == "" {
		b.errs = append(b.errs, ErrLastNameIsRequired)
	}
	if b.firstName == "" {
		b.errs = append(b.errs, ErrFirstNameIsRequired)
	}
	if b.createdDateTime == (time.Time{}.UTC()) {
		b.errs = append(b.errs, ErrCreatedDateTimeProfileIsRequired)
	}
}
