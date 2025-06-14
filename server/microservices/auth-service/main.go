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
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"os/signal"
	"syscall"

	"github.com/QUDUSKUNLE/microservices/auth-service/core/services"
	"github.com/QUDUSKUNLE/microservices/auth-service/adapters/handler"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {
	// Load configuration
	// Load environment variable
	cfg, err := utils.LoadConfig("AUTH")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// Initialize database connection
	db := db.DatabaseConnection(cfg.DB_URL)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Create TCP listener
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("Error starting auth service: %v", err)
	}

	defer listen.Close()

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	// Create gRPC server with TLS and interceptors
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.ValidationInterceptor(),
		))

	// Initialize use case and register services
	authUseCase := services.InitAuthServer(db)

	handler.NewAuthServer(grpcServer, authUseCase)
	reflection.Register(grpcServer)

	go func() {
		logger.GetLogger().Info("Auth Service listening at with TLS enabled (Min version: TLS 1.2)", zap.String("address", cfg.Port))
		if err := grpcServer.Serve(listen); err != nil {
			logger.GetLogger().Fatal("failed to serve auth service", zap.Error(err))
		}
	}()
	<-ctx.Done()
	logger.GetLogger().Info("Shutting down auth service...")
	grpcServer.GracefulStop()

	logger.GetLogger().Info("Auth service stopped gracefully")
}
