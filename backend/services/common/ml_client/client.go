package mlclient

import (
	"context"
	"fmt"

	"github.com/samiransarii/inboXpert/common/utils"
	pb "github.com/samiransarii/inboXpert/services/common/ml_server_protogen"
)

func NewClient(cfg ClientConfig) (*MLPredictionClient, error) {
	manager := utils.GetGRPCClientManager()

	_, err := manager.GetConnection(context.Background(), cfg.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to establish initial connection: %w", err)
	}

	return &MLPredictionClient{
		serviceAddr: cfg.Address,
		manager:     manager,
	}, nil
}

func (c *MLPredictionClient) CategorizeEmail(ctx context.Context, email *pb.EmailRequest) (*pb.CategoryResponse, error) {
	conn, err := c.manager.GetConnection(ctx, c.serviceAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}

	client := pb.NewEmailPredictionClient(conn)
	return client.CategorizeEmail(ctx, email)
}

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

func (c *MLPredictionClient) Close() error {
	return c.manager.CloseConnection(c.serviceAddr)
}
