package mlclient

import (
	"context"
	"fmt"

	"github.com/samiransarii/inboXpert/common/utils"
	pb "github.com/samiransarii/inboXpert/services/common/ml_server_protogen"
)

// NewClient initializes a new MLPredictionClient using the provided configuration.
// It verifies the connection to the ML service during initialization.
func NewClient(cfg ClientConfig) (*MLPredictionClient, error) {
	manager := utils.GetGRPCClientManager()

	// Verify the initial connection to the ML service.
	_, err := manager.GetConnection(context.Background(), cfg.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to establish initial connection: %w", err)
	}

	return &MLPredictionClient{
		serviceAddr: cfg.Address,
		manager:     manager,
	}, nil
}

// CategorizeEmail sends a single email to the ML service for categorization
// and returns the categorization result.
func (c *MLPredictionClient) CategorizeEmail(ctx context.Context, email *pb.EmailRequest) (*pb.CategoryResponse, error) {
	conn, err := c.manager.GetConnection(ctx, c.serviceAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}

	client := pb.NewEmailPredictionClient(conn)
	return client.CategorizeEmail(ctx, email)
}

// BatchCategorizeEmails sends a batch of emails to the ML service for categorization
// and returns a BatchCategoryResponse containing the results for all emails.
func (c *MLPredictionClient) BatchCategorizeEmails(ctx context.Context, emails []*pb.EmailRequest) (*pb.BatchCategoryResponse, error) {
	conn, err := c.manager.GetConnection(ctx, c.serviceAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}

	client := pb.NewEmailPredictionClient(conn)
	request := &pb.BatchEmailRequest{
		Emails: emails,
	}
	return client.BatchCategorizeEmail(ctx, request)
}

// Close terminates the gRPC connection to the ML service.
func (c *MLPredictionClient) Close() error {
	return c.manager.CloseConnection(c.serviceAddr)
}
