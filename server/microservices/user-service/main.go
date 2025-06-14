package main

// Package main is the entry point for the user-service microservice.
// It initializes and starts the gRPC server for handling user-related operations.
//
// This file imports the following packages:
// - "fmt", "log", "net", and "os" from the standard library for basic I/O, logging, networking, and environment variable handling.
// - "github.com/QUDUSKUNLE/microservices/user-service/pkg/v1/handler" for handling user-related gRPC requests.
// - "github.com/QUDUSKUNLE/microservices/user-service/pkg/v1/usercase" for implementing business logic related to user operations.
// - "github.com/QUDUSKUNLE/microservices/events-service/publish" for publishing events to the events-service.
// - "github.com/QUDUSKUNLE/microservices/shared/db" for database connection and operations.
// - "github.com/QUDUSKUNLE/microservices/shared/middleware" for shared middleware functionalities.
// - "github.com/QUDUSKUNLE/microservices/shared/utils" for utility functions shared across microservices.
// - "github.com/segmentio/kafka-go" for interacting with Kafka for message publishing and consumption.
// - "google.golang.org/grpc" and "google.golang.org/grpc/reflection" for setting up and managing the gRPC server.
import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/QUDUSKUNLE/microservices/events-service/publish"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"github.com/QUDUSKUNLE/microservices/user-service/core/services"
	"github.com/QUDUSKUNLE/microservices/user-service/adapters/handler"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	// Load configuration
	// Load environment variable
	cfg, err := utils.LoadConfig("USER")
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
		log.Fatalf("Error starting user service: %v", err)
	}
	defer listen.Close()

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	// Create gRPC server with TLS and interceptors
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.ValidationInterceptor(),
		),
	)
	// Initialize use case and register services
	userUseCase := services.InitUserServer(db)
	eventBroker := publish.NewBroker(cfg.KafkaBroker, cfg.KafkaTopic)
	handler.NewUserService(grpcServer, userUseCase, eventBroker)
	reflection.Register(grpcServer)

	go func() {
		logger.GetLogger().Info("User Service listening at with TLS enabled (Min version: TLS 1.2)", zap.String("address", cfg.Port))
		if err := grpcServer.Serve(listen); err != nil {
			logger.GetLogger().Fatal("failed to serve user service", zap.Error(err))
		}
	}()
	<-ctx.Done()
	// Graceful shutdown
	logger.GetLogger().Info("Shutting down user service")
	grpcServer.GracefulStop()

	logger.GetLogger().Info("User service stopped gracefully")
}
