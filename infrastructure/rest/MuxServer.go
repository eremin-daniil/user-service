package rest

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	jwtServiceInterface "user_service/infrastructure/jwt/interface"
	loggerInterface "user_service/infrastructure/logger/interface"
	serverInterface "user_service/infrastructure/rest/interface"
	"user_service/infrastructure/rest/middleware"
)

type MuxServer struct {
	router     *mux.Router
	server     *http.Server
	logger     loggerInterface.Logger
	jwtService jwtServiceInterface.JWTService
}

func NewMuxServer(logger loggerInterface.Logger, jwtService jwtServiceInterface.JWTService) serverInterface.Server {
	router := mux.NewRouter()

	router.Use(middleware.RequestMiddleware(logger))

	server := &http.Server{Handler: router}
	return &MuxServer{router: router, server: server, logger: logger, jwtService: jwtService}
}

func (s *MuxServer) RegisterPublicRoute(method, path string, handler http.HandlerFunc) {
	s.router.Handle(path, handler).Methods(method)
}

func (s *MuxServer) RegisterPrivateRoute(method, path string, handler http.HandlerFunc) {
	jwtMiddleware := middleware.JWTMiddleware(s.logger, s.jwtService)
	s.router.Handle(path, jwtMiddleware(handler)).Methods(method)
}

func (s *MuxServer) Start(ctx context.Context, address string) {
	s.server.Addr = address
	s.logger.Info(context.Background(), fmt.Sprintf("starting REST server at address: %s", address))
	err := s.server.ListenAndServe()
	if err != nil {
		s.logger.Error(ctx, fmt.Errorf("REST server start failed: %s", err))
	}
}
