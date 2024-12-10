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
	mlClient     mlclient.Service
	grpcServer   *grpc.Server
	categHandler *handlers.CategorizationHandler
	emailRepo    *handlers.EmailRepository
}

func NewServer(config *models.Config) (*Server, error) {
	// Initialize mlClient
	mlClient, err := mlclient.NewClient(mlclient.ClientConfig{
		Address: config.MLServerAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create ML client: %w", err)
	}

	// Initialize the email repository
	emailRepo := handlers.NewEmailRepository(config.DBPool)

	handler := handlers.NewCategorizationHandler(mlClient, config, emailRepo)

	grpcServer := grpc.NewServer()
	pb.RegisterEmailCategorizationServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	return &Server{
		config:       config,
		mlClient:     mlClient,
		grpcServer:   grpcServer,
		categHandler: handler,
		emailRepo:    emailRepo,
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
