package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samiransarii/inboXpert/services/email-categorization/internal/config"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/server"
)

func main() {
	cfg := config.New()

	// Create a new gRPC server instance based on the provided configuration.
	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Start a separate goroutine to handle graceful shutdown.
	// It waits for interrupt signals (e.g., Ctrl+C or SIGTERM) and then
	// stops the gRPC server cleanly.
	go func() {
		sigChan := make(chan os.Signal, 1)

		// Register for notification on the specified signals.
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan // Block until a signal is received.

		log.Println("Shutting down gRPC server...")
		srv.Stop()
	}()

	// Start the gRPC server. If it fails to start or encounters an error,
	// log it and exit.
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
