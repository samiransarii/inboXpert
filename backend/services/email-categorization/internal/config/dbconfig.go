package config

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/samiransarii/inboXpert/common/utils"
)

func connectDB() *pgxpool.Pool {
	// Load environmental variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get the database URL from environment variables
	databaseURL := utils.GetEnv("DATABASE_URL", "")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// parse the connection config
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	// Set additional pool configurations
	config.MaxConns = 10
	config.HealthCheckPeriod = 2 * time.Minute

	// Create a new connection pool
	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Printf("Connected to PostgreSQL!")
	return dbpool
}
