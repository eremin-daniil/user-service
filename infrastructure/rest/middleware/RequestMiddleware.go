package middleware

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"user_service/infrastructure/constant"
	loggerInterface "user_service/infrastructure/logger/interface"
)

func RequestMiddleware(logger loggerInterface.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r, err := appendRequestID(r)
			if err != nil {
				logger.Warning(r.Context(), err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			r = appendPathVariable(r)
			next.ServeHTTP(w, r)
		})
	}
}

func appendRequestID(r *http.Request) (*http.Request, error) {
	header := r.Header.Get("X-Request-ID")
	if header == "" {
		ctx := context.WithValue(r.Context(), constant.RequestIDCtxKey, uuid.New().String())
		return r.WithContext(ctx), nil
	}
	requestID, err := uuid.Parse(header)
	if err != nil {
		return r, fmt.Errorf("invalid request ID: %s", header)
	}
	ctx := context.WithValue(r.Context(), constant.RequestIDCtxKey, requestID.String())
	return r.WithContext(ctx), nil
}

func appendPathVariable(r *http.Request) *http.Request {
	ctx := context.WithValue(r.Context(), constant.PathVariablesCtxKey, mux.Vars(r))
	return r.WithContext(ctx)
}
