package models

import "github.com/jackc/pgx/v5/pgxpool"

type Config struct {
	GRPCPort      string
	MLServerAddr  string
	MaxBatchSize  int
	NumWorkers    int
	RetryAttempts int
	DBPool        *pgxpool.Pool
}
