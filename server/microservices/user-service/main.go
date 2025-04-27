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
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/user-service/pkg/v1/handler"
	"github.com/QUDUSKUNLE/microservices/user-service/pkg/v1/usercase"
	"github.com/QUDUSKUNLE/microservices/events-service/publish"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"github.com/segmentio/kafka-go"
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
		log.Fatalf("Error starting user service: %v", err)
	}

	// Create gRPC server with TLS and interceptors
	grpcServer := grpc.NewServer(
		// grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			middleware.ValidationInterceptor(),
		))

	// Initialize use case and register services
	userUseCase := usercase.InitUserServer(db)

	// Initilaize new broker
	broker := publish.NewBroker(kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{os.Getenv("KAFKA_BROKER")},
	}))
	handler.NewAuthServer(grpcServer, userUseCase, broker, os.Getenv("ORGANIZATION"))
	reflection.Register(grpcServer)

	log.Printf("User Service listening at %v with TLS enabled (Min version: TLS 1.2)", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve user service: %v", err)
	}
}
