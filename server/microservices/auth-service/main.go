package main

// Package main is the entry point for the auth-service microservice.
// It initializes and starts the gRPC server for handling authentication-related requests.
//
// - External dependencies:
//   - github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/authcase: Contains the use case logic for authentication.
//   - github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/handler: Provides gRPC handlers for authentication services.
//   - github.com/QUDUSKUNLE/microservices/shared/db: Provides shared database utilities.
//   - github.com/QUDUSKUNLE/microservices/shared/middleware: Provides shared middleware utilities.
//   - github.com/QUDUSKUNLE/microservices/shared/utils: Provides shared utility functions.
//   - google.golang.org/grpc: Provides gRPC server and client functionality.
//   - google.golang.org/grpc/reflection: Provides server reflection for gRPC debugging and tooling.
import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/authcase"
	"github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/handler"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	// Load environment variable
	if err := utils.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Initialize database connection
	db := db.DatabaseConnection()

	// Create TCP listener
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("Error starting auth service: %v", err)
	}

	// Create gRPC server with TLS and interceptors
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.ValidationInterceptor(),
		))

	// Initialize use case and register services
	authUseCase := authcase.InitAuthServer(db)

	handler.NewAuthServer(grpcServer, authUseCase)
	reflection.Register(grpcServer)

	log.Printf("Auth Service listening at %v with TLS enabled (Min version: TLS 1.2)", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve auth service: %v", err)
	}
}
