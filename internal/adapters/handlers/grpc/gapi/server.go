package gapi

import (
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/config"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/services"
)

type Server struct {
	generated.UnimplementedUserGrpcServiceServer
	user_service services.UserService
	config       config.Config
}

func NewServer(user_service services.UserService, config config.Config) *Server {
	return &Server{
		user_service: user_service,
		config:       config,
	}
}
