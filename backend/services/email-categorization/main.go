package main

import "fmt"

func main() {
	fmt.Printf("Email Categorization Server is up!!")
}

// import (
// 	"context"
// 	"log"
// 	"net"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"

// 	mlclient "github.com/samiransarii/inboXpert/services/common/ml_client"
// 	mlpb "github.com/samiransarii/inboXpert/services/common/ml_server_protogen"
// 	pb "github.com/samiransarii/inboXpert/services/email-categorization/proto"
// )

// type emailCategorizationServer struct {
// 	pb.UnimplementedEmailCategorizationServiceServer
// 	mlClient mlclient.Service
// }

// func newServer() (*emailCategorizationServer, error) {
// 	client, err := mlclient.NewClient(mlclient.ClientConfig{
// 		Address: "localhost:50055",
// 	})
// 	if err != nil {
// 		log.Printf("Failed to create ML client: %v", err)
// 	}

// 	return &emailCategorizationServer{
// 		mlClient: client,
// 	}, nil
// }

// func (s *emailCategorizationServer) CategorizeEmail(ctx context.Context, req *pb.CategorizeRequest) (*pb.CategorizeResponse, error) {
// 	var mlEmails []*mlpb.EmailRequest

// 	mlEmail := &mlpb.EmailRequest{
// 		Id:         req.Email.Id,
// 		Subject:    req.Email.Subject,
// 		Body:       req.Email.Body,
// 		Sender:     req.Email.Sender,
// 		Recipients: req.Email.Recipients,
// 		Headers:    req.Email.Headers,
// 	}
// 	mlEmails = append(mlEmails, mlEmail)

// 	mlResponse, err := s.mlClient.CategorizeEmail(ctx, mlEmails[0])
// 	if err != nil {
// 		log.Printf("Error getting ML prediction: %v", err)
// 	}

// 	result := &pb.CategoryResult{
// 		Id:              mlResponse.Id,
// 		Categories:      []string{mlResponse.Category},
// 		ConfidenceScore: mlResponse.Confidence,
// 	}

// 	for _, alt := range mlResponse.Alternatives {
// 		result.Categories = append(result.Categories, alt.Category)
// 	}

// 	return &pb.CategorizeResponse{Result: result}, nil
// }

// func main() {
// 	server, err := newServer()
// 	if err != nil {
// 		log.Fatalf("Failed to create server: %v", err)
// 	}

// 	listener, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	pb.RegisterEmailCategorizationServiceServer(grpcServer, server)
// 	reflection.Register(grpcServer)

// 	// Handle shutdown gracefully
// 	go func() {
// 		c := make(chan os.Signal, 1)
// 		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
// 		<-c
// 		log.Println("Shutting down gRPC server...")

// 		// Close ML client
// 		if server.mlClient != nil {
// 			if err := server.mlClient.Close(); err != nil {
// 				log.Printf("Error closing ML client: %v", err)
// 			}
// 		}

// 		grpcServer.GracefulStop()
// 	}()

// 	log.Println("gRPC server is litening on port: 50051")
// 	if err := grpcServer.Serve(listener); err != nil {
// 		log.Fatalf("Failed to serve: %v", err)
// 	}
// }
