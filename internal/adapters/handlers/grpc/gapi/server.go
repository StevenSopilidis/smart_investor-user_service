package gapi

import (
	"log"
	"net"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/config"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/services"
	"google.golang.org/grpc"
)

type Server struct {
	generated.UnimplementedUserGrpcServiceServer
	user_service *services.UserService
	config       config.Config
}

func NewServer(user_service *services.UserService, config config.Config) *Server {
	return &Server{
		user_service: user_service,
		config:       config,
	}
}

func (s *Server) Run() {
	grpc := grpc.NewServer()
	generated.RegisterUserGrpcServiceServer(grpc, s)

	listener, err := net.Listen("tcp", s.config.GRPCAddress)
	if err != nil {
		log.Fatal("---> Could not create listener: ", err)
	}

	log.Println("---> Server starting listening at: ", s.config.GRPCAddress)
	err = grpc.Serve(listener)
	if err != nil {
		log.Fatalf("Could not listen at address: %s due to Error: %s", s.config.GRPCAddress, err)
	}
}
