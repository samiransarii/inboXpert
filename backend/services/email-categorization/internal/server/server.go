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

// Server initializes and runs a gRPC server for email categorization.
// It sets up the ML client, email repository, and categorization handler,
// then registers the gRPC service and manages startup/shutdown.
type Server struct {
	config       *models.Config
	mlClient     mlclient.Service
	grpcServer   *grpc.Server
	categHandler *handlers.CategorizationHandler
	emailRepo    *handlers.EmailRepository
}

// NewServer creates a new Server instance, configuring the ML client, repository, handlers,
// and the gRPC server. It returns an error if any of the components fail to initialize.
func NewServer(config *models.Config) (*Server, error) {
	// Initialize the machine learning client
	mlClient, err := mlclient.NewClient(mlclient.ClientConfig{
		Address: config.MLServerAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create ML client: %w", err)
	}

	// Initialize the email repository for database operations
	emailRepo := handlers.NewEmailRepository(config.DBPool)

	// Create the categorization handler that ties everything together
	handler := handlers.NewCategorizationHandler(mlClient, config, emailRepo)

	// Create and register the gRPC server and reflection service
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

// Start begins listening on the configured gRPC port and handles incoming requests.
// If the server fails to start listening, it returns an error.
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.config.GRPCPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	log.Printf("gRPC server is listening on port: %s", s.config.GRPCPort)
	return s.grpcServer.Serve(listener)
}

// Stop gracefully stops the gRPC server and closes the ML client, ensuring no new
// requests are accepted and ongoing requests are completed before shutdown.
func (s *Server) Stop() {
	if s.mlClient != nil {
		if err := s.mlClient.Close(); err != nil {
			log.Printf("Error closing ML client: %v", err)
		}
	}
	s.grpcServer.GracefulStop()
}
