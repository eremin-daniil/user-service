package user

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	userEntity "user_service/domain/entity/user"
	"user_service/domain/entity/user/primitive"
	profileEntity "user_service/domain/entity/user/profile"
)

type BuilderShould struct {
	suite.Suite
}

func TestBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(BuilderShould))
}

func (s *BuilderShould) TestBuilderBuild_ValidValues_ReturnUser() {
	expectedID := uuid.NewString()
	expectedLogin := "Login"
	expectedProfile := &profileEntity.Profile{}

	actual, err := userEntity.NewBuilder().
		IDString(expectedID).
		LoginString(expectedLogin).
		Profile(expectedProfile).
		Build()

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedID, actual.Id().String())
	assert.Equal(s.T(), expectedLogin, actual.Login().String())
	assert.Equal(s.T(), expectedProfile, actual.Profile())
}

func (s *BuilderShould) TestBuilderBuild_InvalidValues_ReturnError() {
	expectedErr := errors.Join(
		userEntity.ErrInvalidUserID,
		primitive.ErrIncorrectLengthLogin,
		userEntity.ErrLoginIsRequired,
	)

	actual, err := userEntity.NewBuilder().
		IDString("").
		LoginString("").
		Profile(nil).
		Build()

	assert.Nil(s.T(), actual)
	assert.Equal(s.T(), expectedErr, err)
}

func (s *BuilderShould) TestBuilderBuild_NotFillingRequiredFields_ReturnError() {
	expectedErr := errors.Join(
		userEntity.ErrInvalidUserID,
		userEntity.ErrLoginIsRequired,
	)

	actual, err := userEntity.NewBuilder().
		IDString("").
		Profile(nil).
		Build()

	assert.Nil(s.T(), actual)
	assert.Equal(s.T(), expectedErr, err)
}
