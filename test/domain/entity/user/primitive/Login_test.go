package primitive

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	userPrimitive "user_service/domain/entity/user/primitive"
)

type LoginShould struct {
	suite.Suite
}

func TestLoginShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(LoginShould))
}

func (s *LoginShould) TestLoginFromString_ValidValue_ReturnLogin() {
	expected := "login_success"

	actual, err := userPrimitive.LoginFromString(expected)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expected, actual.String())
}

func (s *LoginShould) TestLoginFromString_IncorrectLengthValue_ReturnError() {
	testCases := []string{
		"0123",
		"0123456789abcdefg",
	}

	for index := range testCases {
		actual, err := userPrimitive.LoginFromString(testCases[index])

		assert.Empty(s.T(), actual)
		assert.Equal(s.T(), userPrimitive.ErrIncorrectLengthLogin, err)
	}
}

func (s *LoginShould) TestLoginFromString_InvalidCharacters_ReturnError() {
	testCases := []string{
		"0123.",
		"0123-",
		"0123+",
		"0123/",
		"0123*",
		"0123[",
		"0123]",
	}

	for index := range testCases {
		actual, err := userPrimitive.LoginFromString(testCases[index])

		assert.Empty(s.T(), actual)
		assert.Equal(s.T(), userPrimitive.ErrInvalidCharactersLogin, err)
	}
}
