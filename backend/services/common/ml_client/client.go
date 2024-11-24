package mlclient

import (
	"sync"
	"time"

	pb "github.com/samiransarii/inboXpert/services/common/ml_server_protogen"
	"google.golang.org/grpc"
)

type MLPredictionClient struct {
	pClient     pb.EmailPredictionClient
	conn        *grpc.ClientConn
	mu          sync.Mutex
	serviceAddr string
	isConnected bool
}

type ClientConfig struct {
	Address     string
	DialTimeout time.Duration
}


