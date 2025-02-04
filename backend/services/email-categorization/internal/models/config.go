package models

import "github.com/jackc/pgx/v5/pgxpool"

// Config holds configuration data for the email categorization service.
// This includes server settings, connection details for the ML service,
// processing parameters, retry policies, and a database connection pool.
type Config struct {
	GRPCPort      string       // gRPC server port
	MLServerAddr  string       // ML service address
	MaxBatchSize  int          // Max emails in a single batch
	NumWorkers    int          // Concurrent worker count
	RetryAttempts int          // Retry count for failed operations
	DBPool        *pgxpool.Pool
}
