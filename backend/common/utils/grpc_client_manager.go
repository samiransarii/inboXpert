package utils

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type GRPCClientManager struct {
	connections map[string]*grpc.ClientConn
	mutex       sync.RWMutex
}

var (
	manager *GRPCClientManager
	once    sync.Once
)

func GetGRPCClientManager() *GRPCClientManager {
	once.Do(func() {
		manager = &GRPCClientManager{
			connections: make(map[string]*grpc.ClientConn),
		}
	})
	return manager
}

func (m *GRPCClientManager) GetConnection(ctx context.Context, serviceAddr string) (*grpc.ClientConn, error) {
	// Try to get existing connection
	m.mutex.RLock()
	if conn, exists := m.connections[serviceAddr]; exists {
		m.mutex.RUnlock()
		return conn, nil
	}
	m.mutex.RUnlock()

	// Create new connection if none exists
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// Double check after acquiring write lock
	if conn, exists := m.connections[serviceAddr]; exists {
		return conn, nil
	}

	// Configure client options
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             3 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.WithDefaultServiceConfig(`{
            "methodConfig": [{
                "name": [{"service": ""}],
                "retryPolicy": {
                    "maxAttempts": 3,
                    "initialBackoff": "0.1s",
                    "maxBackoff": "1s",
                    "backoffMultiplier": 2.0,
                    "retryableStatusCodes": ["UNAVAILABLE"]
                }
            }]
        }`),
	}

	// Create new connection
	conn, err := grpc.NewClient(serviceAddr, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}

	// Store the connection
	m.connections[serviceAddr] = conn

	return conn, nil
}

func (m *GRPCClientManager) CloseConnection(serviceAddr string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if conn, exists := m.connections[serviceAddr]; exists {
		if err := conn.Close(); err != nil {
			return fmt.Errorf("failed to close connection: %w", err)
		}
		delete(m.connections, serviceAddr)
	}
	return nil
}

func (m *GRPCClientManager) CloseAllConnections() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var lastErr error
	for addr, conn := range m.connections {
		if err := conn.Close(); err != nil {
			lastErr = fmt.Errorf("failed to close connection to %s: %w", addr, err)
			log.Printf("Error closing connection to %s: %v", addr, err)
		}
		delete(m.connections, addr)
	}
	return lastErr
}
