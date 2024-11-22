package main

import (
	"context"
	"log"
	"net"

	pb "github.com/samiransarii/inboXpert/backend/proto/gen/services/email_categorization_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type emailCategorizationServer struct {
	pb.UnimplementedEmailCategorizationServiceServer
}

func (s *emailCategorizationServer) CategorizeEmails(ctx context.Context, req *pb.CategorizeRequest) (*pb.CategorizeResponse, error) {
	var results []*pb.CategoryResult

	for _, email := range req.Emails {
		// Simple categorization logic

		categories := []string{}
		confidenceScore := float32(0.9)

		if email.Subject == "" && email.Body == "" {
			categories = append(categories, "Empty")
			confidenceScore = float32(0.5)
		} else {
			categories = append(categories, "General")
		}

		result := &pb.CategoryResult{
			EmailId:         email.Id,
			Categories:      categories,
			ConfidenceScore: confidenceScore,
			Error:           "",
		}

		results = append(results, result)
	}

	return &pb.CategorizeResponse{Results: results}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEmailCategorizationServiceServer(grpcServer, &emailCategorizationServer{})

	reflection.Register(grpcServer)

	log.Println("gRPC server is litening on port: 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
