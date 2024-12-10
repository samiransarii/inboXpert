package config

import "github.com/samiransarii/inboXpert/services/email-categorization/internal/models"

func New() *models.Config {
	return &models.Config{
		GRPCPort:      ":50051",
		MLServerAddr:  "localhost:50055",
		MaxBatchSize:  1000,
		NumWorkers:    10,
		RetryAttempts: 3,
	}
}

