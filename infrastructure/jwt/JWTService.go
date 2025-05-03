package jwtService

import (
	"context"
	"github.com/google/uuid"
	"time"
	"user_service/infrastructure/constant"
)
import "github.com/golang-jwt/jwt/v4"

const (
	defaultJwtTokenTTL = 60 * 60 * 24 * 2 // 2 day
)

type JWTService struct {
	secret      string
	claims      map[string]bool
	jwtTokenTTL time.Duration
}

func NewJWTService(secret string, jwtTokenTTL time.Duration) *JWTService {
	claims := map[string]bool{
		roleTokenKey:       true,
		scopeTokenKey:      true,
		userIDTokenKey:     true,
		expirationTokenKey: true,
	}
	if jwtTokenTTL == 0 {
		jwtTokenTTL = defaultJwtTokenTTL
	}
	return &JWTService{secret: secret, claims: claims, jwtTokenTTL: jwtTokenTTL}
}

func (j *JWTService) Verify(tokenString string) bool {
	_, err := j.parse(tokenString)
	return err == nil
}

func (j *JWTService) FillCtxWithParams(ctx context.Context, tokenString string) (context.Context, error) {
	token, err := j.parse(tokenString)
	if err != nil {
		return ctx, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx, ErrJWTWrongClaims
	}
	userID, err := j.extractClaimAsString(claims, userIDTokenKey)
	if err != nil {
		return ctx, err
	}
	_, err = uuid.Parse(userID)
	if err != nil {
		return ctx, ErrJWTInvalidUserID
	}
	role, err := j.extractClaimAsString(claims, roleTokenKey)
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(ctx, constant.UserIDCtxKey, userID)
	ctx = context.WithValue(ctx, constant.UserRoleCtxKey, role)

	return ctx, nil
}

func (j *JWTService) CreateUserToken(userID string, claims map[string]string) (string, error) {
	if userID == "" {
		return "", ErrJWTMissingUserID
	}
	expirationTime := time.Now().Add(j.jwtTokenTTL * time.Second).Unix()
	tokenClaims := jwt.MapClaims{
		userIDTokenKey:     userID,
		expirationTokenKey: expirationTime,
	}
	for key, value := range claims {
		if !j.claims[key] {
			return "", ErrJWTInvalidClaims
		}
		tokenClaims[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	signedToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", ErrSignJWT
	}
	return signedToken, nil

}

func (j *JWTService) parse(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		switch token.Method.(type) {
		case *jwt.SigningMethodHMAC:
			return []byte(j.secret), nil
		default:
			return nil, ErrJWTUnsupportedSigningMethod
		}
	})
}

func (j *JWTService) extractClaimAsString(claims jwt.MapClaims, key string) (string, error) {
	value, ok := claims[key]
	if !ok {
		return "", ErrJWTMissingClaim
	}
	strValue, ok := value.(string)
	if !ok {
		return "", ErrJWTInvalidClaimType
	}
	return strValue, nil
}
