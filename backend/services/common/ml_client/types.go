package mlclient

import (
	"context"

	"github.com/samiransarii/inboXpert/common/utils"
	pb "github.com/samiransarii/inboXpert/services/common/ml_server_protogen"
)

// MLPredictionClient implements the Service interface and provides methods
// to interact with the ML service for email categorization.
type MLPredictionClient struct {
	serviceAddr string
	manager     *utils.GRPCClientManager
}

// ClientConfig holds the configuration required to connect to the ML service.
type ClientConfig struct {
	Address string
}

// Service defines the interface for interacting with the ML service.
// It includes methods for single and batch email categorization and for closing the connection.
// @Depricated
type Service interface {
	// CategorizeEmail sends a single email to the ML service for categorization
	// and returns the category and confidence score.
	CategorizeEmail(ctx context.Context, email *pb.EmailRequest) (*pb.CategoryResponse, error)

	// BatchCategorizeEmails sends a batch of emails to the ML service for categorization
	// and returns the results for all emails in the batch.
	BatchCategorizeEmails(ctx context.Context, emails []*pb.EmailRequest) (*pb.BatchCategoryResponse, error)

	// Close shuts down the connection to the ML service.
	Close() error
}
