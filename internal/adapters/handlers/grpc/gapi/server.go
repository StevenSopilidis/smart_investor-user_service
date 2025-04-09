package gapi

import (
	"context"
	"log"
	"net"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/config"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/observability"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/services"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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
	// initialize tracer provider
	ctx1 := context.Background()
	trace_cleanup, err := observability.InitTracer(ctx1, s.config)
	if err != nil {
		log.Fatal("---> Could not initialize tracer provider: ", err)
	}
	defer trace_cleanup()

	// initialize meter provider
	ctx2 := context.Background()
	metric_cleanup, err := observability.InitMeterProvider(ctx2, s.config)
	if err != nil {
		log.Fatal("---> Could not initilize meter provider: ", err)
	}
	defer metric_cleanup(context.Background())

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	generated.RegisterUserGrpcServiceServer(grpcServer, s)

	listener, err := net.Listen("tcp", s.config.GRPCAddress)
	if err != nil {
		log.Fatal("---> Could not create listener: ", err)
	}

	log.Println("---> Server starting listening at: ", s.config.GRPCAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Could not listen at address: %s due to Error: %s", s.config.GRPCAddress, err)
	}
}
