package profile

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
	profileEntity "user_service/domain/entity/user/profile"
)

type ProfileShould struct {
	suite.Suite
}

func TestProfileShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProfileShould))
}

func (s *ProfileShould) TestProfileChangeFirstAndLastName_ValidLastAndFirstName_Return() {
	expectedLastName := "Name"
	expectedFirstName := "Name"
	profile := profileEntity.Profile{}

	err := profile.ChangeFirstAndLastName(expectedLastName, expectedFirstName)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedLastName, profile.LastName())
	assert.Equal(s.T(), expectedFirstName, profile.FirstName())
	assert.NotEqual(s.T(), time.Time{}.UTC(), profile.UpdatedDateTime())
}

func (s *ProfileShould) TestProfileChangeFirstAndLastName_IncorrectLengthLastName_ReturnError() {
	testCases := []string{
		"",
		"A",
		"Abcdefghijklmnopqrstuvwxyzabcdefg",
	}

	for index := range testCases {
		profile := profileEntity.Profile{}

		err := profile.ChangeFirstAndLastName(testCases[index], "Name")

		assert.Empty(s.T(), profile.LastName())
		assert.Empty(s.T(), profile.FirstName())
		assert.Equal(s.T(), time.Time{}.UTC(), profile.UpdatedDateTime())
		assert.Equal(s.T(), profileEntity.ErrIncorrectLengthLastName, err)
	}
}

func (s *ProfileShould) TestProfileChangeFirstAndLastName_IncorrectLengthFirstName_ReturnError() {
	testCases := []string{
		"",
		"A",
		"Abcdefghijklmnopqrstuvwxyzabcdefg",
	}

	for index := range testCases {
		profile := profileEntity.Profile{}

		err := profile.ChangeFirstAndLastName("Name", testCases[index])

		assert.Empty(s.T(), profile.LastName())
		assert.Empty(s.T(), profile.FirstName())
		assert.Equal(s.T(), time.Time{}.UTC(), profile.UpdatedDateTime())
		assert.Equal(s.T(), profileEntity.ErrIncorrectLengthFirstName, err)
	}
}
