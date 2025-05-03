package _interface

import "context"

type JWTService interface {
	Verify(tokenString string) bool
	FillCtxWithParams(ctx context.Context, tokenString string) (context.Context, error)
	CreateUserToken(userID string, claims map[string]string) (string, error)
}
