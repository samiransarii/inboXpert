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

// GRPCClientManager manages the lifecycle of multiple gRPC client connections,
// allowing reuse of connections and ensuring thread-safe access. It provides
// methods to retrieve, create, and close connections as needed.
type GRPCClientManager struct {
	connections map[string]*grpc.ClientConn
	mutex       sync.RWMutex
}

var (
	manager *GRPCClientManager
	once    sync.Once
)

// GetGRPCClientManager returns a singleton instance of GRPCClientManager.
// Using sync.Once ensures that this instance is initialized only once,
// which helps maintain a global connection manager throughout the application.
func GetGRPCClientManager() *GRPCClientManager {
	once.Do(func() {
		manager = &GRPCClientManager{
			connections: make(map[string]*grpc.ClientConn),
		}
	})
	return manager
}

// GetConnection retrieves a gRPC client connection for the provided serviceAddr.
// If a connection for the given address already exists, it returns that connection.
// Otherwise, it establishes a new connection with configured options, stores it,
// and returns it.
//
// It uses a read-write lock to safely manage concurrent access to the connection map.
// If the connection does not exist initially, the manager acquires a write lock and
// creates a new connection. Subsequent calls for the same address will return the
// cached connection.
func (m *GRPCClientManager) GetConnection(ctx context.Context, serviceAddr string) (*grpc.ClientConn, error) {
	// First, attempt to retrieve the connection using a read lock.
	m.mutex.RLock()
	if conn, exists := m.connections[serviceAddr]; exists {
		m.mutex.RUnlock()
		return conn, nil
	}
	m.mutex.RUnlock()

	// Acquire write lock to create a new connection if needed.
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// Double-check if another goroutine created the connection before we acquired the lock.
	if conn, exists := m.connections[serviceAddr]; exists {
		return conn, nil
	}

	// Set up connection options:
	// - Insecure credentials for transport (suitable for local dev or if using an external TLS layer).
	// - Keepalive parameters to maintain healthy connections.
	// - Default service config enabling retries on transient errors (UNAVAILABLE).
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

	// Create a new gRPC client connection.
	conn, err := grpc.NewClient(serviceAddr, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}

	// Store the newly created connection in the manager.
	m.connections[serviceAddr] = conn

	return conn, nil
}

// CloseConnection closes a specific gRPC client connection identified by serviceAddr,
// if it exists. Once closed, it removes the connection from the manager's internal map.
// If the connection does not exist, it does nothing.
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

// CloseAllConnections gracefully closes all open gRPC client connections managed
// by the GRPCClientManager. It returns the last encountered error, if any.
//
// This method can be used during application shutdown to ensure that all resources
// are released properly.
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
