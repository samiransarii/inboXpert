package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/samiransarii/inboXpert/services/email-categorization/internal/handlers"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models"

	mlclient "github.com/samiransarii/inboXpert/services/common/ml_client"
	pb "github.com/samiransarii/inboXpert/services/email-categorization/proto"
)

type Server struct {
	config       *models.Config
	grpcServer   *grpc.Server
	categHandler *handlers.CategorizationHandler
	mlClient     mlclient.Service
}

func NewServer(config *models.Config) (*Server, error) {
	mlClient, err := mlclient.NewClient(mlclient.ClientConfig{
		Address: config.MLServerAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create ML client: %w", err)
	}

	handler := handlers.NewCategorizationHandler(mlClient, config)

	grpcServer := grpc.NewServer()
	pb.RegisterEmailCategorizationServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	return &Server{
		config:       config,
		grpcServer:   grpcServer,
		categHandler: handler,
		mlClient:     mlClient,
	}, nil
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.config.GRPCPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	log.Printf("gRPC server is listening on port: %s", s.config.GRPCPort)
	return s.grpcServer.Serve(listener)
}

func (s *Server) Stop() {
	if s.mlClient != nil {
		if err := s.mlClient.Close(); err != nil {
			log.Printf("Error closing ML client: %v", err)
		}
	}
	s.grpcServer.GracefulStop()
}
