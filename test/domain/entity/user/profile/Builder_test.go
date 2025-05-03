package profile

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
	profileEntity "user_service/domain/entity/user/profile"
)

type BuilderShould struct {
	suite.Suite
}

func TestBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(BuilderShould))
}

func (s *BuilderShould) TestBuilderBuild_ValidValues_ReturnProfile() {
	expectedID := uuid.NewString()
	expectedLastName := "Name"
	expectedFirstName := "Name"
	expectedDateTime := time.Now()

	actual, err := profileEntity.NewBuilder().
		IDString(expectedID).
		LastName(expectedLastName).
		FirstName(expectedFirstName).
		CreatedDateTime(expectedDateTime).
		UpdatedDateTime(expectedDateTime).
		Build()

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedID, actual.Id().String())
	assert.Equal(s.T(), expectedLastName, actual.LastName())
	assert.Equal(s.T(), expectedFirstName, actual.FirstName())
	assert.Equal(s.T(), expectedDateTime, actual.CreatedDateTime())
	assert.Equal(s.T(), expectedDateTime, actual.UpdatedDateTime())
}

func (s *BuilderShould) TestBuilderBuild_InvalidValues_ReturnError() {
	expectedErr := errors.Join(
		profileEntity.ErrInvalidProfileID,
		profileEntity.ErrIncorrectLengthLastName,
		profileEntity.ErrIncorrectLengthFirstName,
		profileEntity.ErrLastNameIsRequired,
		profileEntity.ErrFirstNameIsRequired,
	)

	actual, err := profileEntity.NewBuilder().
		IDString("").
		LastName("N").
		FirstName("N").
		CreatedDateTime(time.Now().UTC()).
		Build()

	assert.Nil(s.T(), actual)
	assert.Equal(s.T(), expectedErr, err)
}

func (s *BuilderShould) TestBuilderBuild_NotFillingRequiredFields_ReturnError() {
	expectedErr := errors.Join(
		profileEntity.ErrLastNameIsRequired,
		profileEntity.ErrFirstNameIsRequired,
		profileEntity.ErrCreatedDateTimeProfileIsRequired,
	)

	actual, err := profileEntity.NewBuilder().
		IDString(uuid.NewString()).
		Build()

	assert.Nil(s.T(), actual)
	assert.Equal(s.T(), expectedErr, err)
}
