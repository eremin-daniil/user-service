package _interface

import (
	"context"
	"net/http"
)

type Server interface {
	RegisterPublicRoute(method, path string, handler http.HandlerFunc)
	RegisterPrivateRoute(method, path string, handler http.HandlerFunc)
	Start(ctx context.Context, address string)
}
