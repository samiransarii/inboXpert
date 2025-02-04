package config

import "github.com/samiransarii/inboXpert/services/email-categorization/internal/models"

// New creates a new Config instance with initialized settings for the email categorization service.
// This includes gRPC server settings, ML server endpoint, batching parameters, worker counts, and retry policies.
// It also establishes a database connection pool for use by other parts of the application.
func New() *models.Config {
	// Create a database connection pool to handle database interactions.
	// The connectDB() function should return a pooled connection object.
	dbPool := connectDB()

	// Return a new Config instance populated with essential parameters.
	return &models.Config{
		// GRPCPort defines the network address and port on which the gRPC server will listen.
		GRPCPort: ":50051",

		// MLServerAddr specifies the address of the machine learning server that handles categorization logic.
		MLServerAddr: "localhost:50055",

		// MaxBatchSize limits how many items can be processed in a batch.
		MaxBatchSize: 1000,

		// NumWorkers sets the number of concurrent workers handling categorization tasks.
		NumWorkers: 10,

		// RetryAttempts determines how many times a failed operation (like DB queries or external calls)
		// will be retried before giving up.
		RetryAttempts: 3,

		// DBPool is the connection pool to the underlying database.
		DBPool: dbPool,
	}
}
