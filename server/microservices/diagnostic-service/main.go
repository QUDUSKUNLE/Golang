package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/diagnostic-service/adapters/repository"
	"github.com/QUDUSKUNLE/microservices/diagnostic-service/adapters/server"
	"github.com/QUDUSKUNLE/microservices/diagnostic-service/adapters/subscribe"
	// "github.com/QUDUSKUNLE/microservices/diagnostic-service/core/services"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"github.com/QUDUSKUNLE/microservices/shared/events"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	// Load environment variable
	if err := utils.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Initialize database connection
	dbase := db.DatabaseConnection()
	port := os.Getenv("PORT")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error starting record service: %v", err)
	}

	defer listen.Close()

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	// Configure the gRPC server with enhanced options
	grpcServer := grpc.NewServer()

	// Subscribe to events
	go events.EventsSubscriber(
		os.Getenv("KAFKA_BROKER"),
		os.Getenv("KAFKA_TOPIC"),
		os.Getenv("KAFKA_GROUP_ID"),
		subscribe.ProcessEvent,
	)
	// Initialize the repository and service
	repo := repository.NewDiagnosticRepository(dbase)
	server.NewDiagnosticServer(grpcServer, repo)
	logger.GetLogger().Info("Diagnostic Service listening on with TLS enabled (Min version: TLS 1.2)", zap.String("address", port))
	if err := grpcServer.Serve(listen); err != nil {
		logger.GetLogger().Fatal("failed to serve record service", zap.Error(err))
	}
}
