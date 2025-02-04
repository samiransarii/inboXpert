package config

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/samiransarii/inboXpert/common/utils"
)

// connectDB sets up and returns a database connection pool using the pgxpool library.
// It reads configuration from environment variables, supports a .env file for local development,
// and applies basic connection pool settings.
//
// 1. Load environment variables from .env if it exists, otherwise rely on system environment variables.
// 2. Retrieve and validate the DATABASE_URL environment variable.
// 3. Parse the database URL into a pgxpool configuration.
// 4. Set pool options like maximum connections and health check intervals.
// 5. Initialize and return a connection pool to the PostgreSQL database.
func connectDB() *pgxpool.Pool {
	// Attempt to load environment variables from .env file, if present.
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Retrieve the database URL from environment variables.
	databaseURL := utils.GetEnv("DATABASE_URL", "")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Parse the database URL to obtain a pgxpool configuration.
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	// Set connection pool parameters, such as the maximum number of connections
	// and how frequently to run health checks.
	config.MaxConns = 10
	config.HealthCheckPeriod = 2 * time.Minute

	// Create the connection pool using the provided configuration.
	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Printf("Connected to PostgreSQL!")
	return dbpool
}
