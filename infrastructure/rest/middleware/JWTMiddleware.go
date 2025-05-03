package middleware

import (
	"net/http"
	"strings"
	jwtInterface "user_service/infrastructure/jwt/interface"
	loggerInterface "user_service/infrastructure/logger/interface"
)

func JWTMiddleware(logger loggerInterface.Logger, jwtService jwtInterface.JWTService) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := getToken(r)
			if token == "" {
				http.Error(w, "missing authorization token", http.StatusUnauthorized)
				return
			}
			if !jwtService.Verify(token) {
				http.Error(w, "invalid authorization token", http.StatusUnauthorized)
				return
			}
			ctx, err := jwtService.FillCtxWithParams(r.Context(), token)
			if err != nil {
				logger.Error(ctx, err)
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			handler.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ")
	}
	return token
}
