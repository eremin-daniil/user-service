package jwtService

import "errors"

var (
	ErrJWTUnsupportedSigningMethod = errors.New("неподдерживаемый метод подписи")
	ErrJWTWrongClaims              = errors.New("неправильный claims")
	ErrJWTMissingClaim             = errors.New("отсутствует claim")
	ErrJWTInvalidClaimType         = errors.New("неверный тип claim")
	ErrJWTInvalidClaims            = errors.New("неверный claims provided")
	ErrJWTMissingUserID            = errors.New("не заполнен userID")
	ErrJWTInvalidUserID            = errors.New("неправильный userID")
	ErrSignJWT                     = errors.New("не удалось подписать токен")
)
