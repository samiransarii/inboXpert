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

	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Handle shutdown gracefully
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		log.Println("Shutting down gRPC server...")
		srv.Stop()
	}()

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
