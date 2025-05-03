package profile

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrIncorrectLengthLastName  = errors.New("incorrect length last name")
	ErrIncorrectLengthFirstName = errors.New("incorrect length first name")
)

type ID struct {
	uuid.UUID
}

type Profile struct {
	id              *ID
	lastName        string
	firstName       string
	createdDateTime time.Time
	updatedDateTime time.Time
}

func (p *Profile) Id() *ID {
	return p.id
}

func (p *Profile) LastName() string {
	return p.lastName
}

func (p *Profile) FirstName() string {
	return p.firstName
}

func (p *Profile) CreatedDateTime() time.Time {
	return p.createdDateTime
}

func (p *Profile) UpdatedDateTime() time.Time {
	return p.updatedDateTime
}

func (p *Profile) ChangeFirstAndLastName(lastName, firstName string) error {
	if len(lastName) < 2 || len(lastName) > 32 {
		return ErrIncorrectLengthLastName
	}
	if len(firstName) < 2 || len(firstName) > 32 {
		return ErrIncorrectLengthFirstName
	}
	p.lastName = lastName
	p.firstName = firstName
	p.updatedDateTime = time.Now().UTC()
	return nil
}
