package mlclient

import (
	"context"

	"github.com/samiransarii/inboXpert/common/utils"
	pb "github.com/samiransarii/inboXpert/services/common/ml_server_protogen"
)

type MLPredictionClient struct {
	serviceAddr string
	manager     *utils.GRPCClientManager
}

type ClientConfig struct {
	Address string
}

type Service interface {
	CategorizeEmail(ctx context.Context, email *pb.EmailRequest) (*pb.CategoryResponse, error)
	BatchCategorizeEmails(ctx context.Context, emails []*pb.EmailRequest) (*pb.BatchCategoryResponse, error)
	Close() error
}
