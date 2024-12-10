package models

type Config struct {
	GRPCPort      string
	MLServerAddr  string
	MaxBatchSize  int
	NumWorkers    int
	RetryAttempts int
}
