package test

import (
	"context"
	"crypto/tls"
	"os"
	"testing"
	"time"

	"github.com/QUDUSKUNLE/microservices/record-service/clients"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

// TestTLSConnection verifies that a client can connect to the TLS-enabled server
func TestTLSConnection(t *testing.T) {
	// Get the server address from environment or use default
	serverAddr := os.Getenv("RECORD_SERVICE_ADDR")
	if serverAddr == "" {
		serverAddr = "localhost:3012" // Default port from your app
	}

	// Test context with timeout to avoid hanging (increased timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Helper function to create TLS credentials for testing
	createTestTLSCredentials := func() (credentials.TransportCredentials, error) {
		// For testing purposes, we'll use InsecureSkipVerify
		// In production, you would properly validate certificates
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true, // Skip certificate verification for testing
			ServerName:         "localhost", // Specify the server name
		}
		return credentials.NewTLS(tlsConfig), nil
	}

	// Test 1: Connect with TLS enabled
	t.Run("Connect with TLS", func(t *testing.T) {
		// Create TLS credentials that handle self-signed certificates
		creds, err := createTestTLSCredentials()
		if err != nil {
			t.Fatalf("Failed to create TLS credentials: %v", err)
		}
		
		t.Logf("Connecting to server at %s with TLS", serverAddr)

		// Attempt to connect with TLS
		conn, err := grpc.DialContext(
			ctx,
			serverAddr,
			grpc.WithTransportCredentials(creds),
			grpc.WithBlock(),
		)
		
		if err != nil {
			// Check if the error is related to connection issues (which might happen if server is not running)
			// rather than TLS configuration problems
			st, ok := status.FromError(err)
			if ok {
				t.Logf("gRPC status error: %s", st.Message())
			}
			t.Fatalf("Failed to connect with TLS: %v", err)
		}
		defer conn.Close()
		
		t.Log("Successfully connected to server with TLS")
	})

	// Test 2: Use the updated client code with TLS
	t.Run("Client with TLS", func(t *testing.T) {
		// Create a custom ClientOptions that uses InsecureSkipVerify for testing
		clientOptions := &clients.ClientOptions{
			UseTLS:       true,
			CertFilePath: "../certs/server.crt",
			// For a real implementation, we would add these options to the client.go file
			// but for testing, we're handling it differently
		}
		
		// We're testing organization client as an example
		// Real implementation would depend on your service structure
		organizationClient := clients.NewGRPClientOrganizationService(
			serverAddr,
			clientOptions,
		)
		
		if organizationClient == nil {
			t.Fatalf("Failed to create organization client with TLS")
		}
		
		t.Log("Successfully created client with TLS")
		
		// Here you would typically call a method on the client
		// to verify the connection works end-to-end
		// For example: organizationClient.SomeMethod(ctx, request)
	})
}

