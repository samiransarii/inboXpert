package main

import (
	pb "github.com/samiransarii/inboXpert/common/ml_server"
)

type emailPredictionServer struct {
	pb.UnimplementedMLServiceServer
}
