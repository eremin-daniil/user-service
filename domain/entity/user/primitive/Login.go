package primitive

import (
	"errors"
	"regexp"
)

var (
	ErrIncorrectLengthLogin   = errors.New("incorrect length login")
	ErrInvalidCharactersLogin = errors.New("invalid characters login")
)

type Login string

func LoginFromString(value string) (Login, error) {
	if len(value) < 5 || len(value) > 16 {
		return "", ErrIncorrectLengthLogin
	} else if regexp.MustCompile("[.\\-+/*\\[\\]]").MatchString(value) {
		return "", ErrInvalidCharactersLogin
	}
	return Login(value), nil
}

func (l Login) String() string {
	return string(l)
}
