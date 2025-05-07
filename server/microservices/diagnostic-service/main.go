package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"context"

	"github.com/QUDUSKUNLE/microservices/diagnostic-service/adapters/repository"
	"github.com/QUDUSKUNLE/microservices/diagnostic-service/adapters/server"
	"github.com/QUDUSKUNLE/microservices/diagnostic-service/adapters/subscribe"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/events"
	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	// Load environment variable
	cfg, err := utils.LoadConfig("DIAGNOSTIC")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// Initialize database connection
	dbase := db.DatabaseConnection(cfg.DB_URL)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start gRPC server
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("Error starting record service: %v", err)
	}
	defer listen.Close()

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	// Configure the gRPC server with enhanced options
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.ValidationInterceptor(),
		),
	)

	// Subscribe to events
	go events.EventsSubscriber(
		cfg.KafkaBroker,
		cfg.KafkaTopic,
		cfg.KafkaGroup,
		subscribe.ProcessEvent,
	)
	// Initialize the repository and service
	repo := repository.NewDiagnosticRepository(dbase)
	server.NewDiagnosticServer(grpcServer, repo)
	go func() {
		logger.GetLogger().Info("Diagnostic Service listening on with TLS enabled (Min version: TLS 1.2)", zap.String("address", cfg.Port))
	if err := grpcServer.Serve(listen); err != nil {
		logger.GetLogger().Fatal("failed to serve record service", zap.Error(err))
	}
}()
<- ctx.Done()
	logger.GetLogger().Info("Shutting down server gracefully...")
	grpcServer.GracefulStop()

	logger.GetLogger().Info("Diagnostic service stopped gracefully")
}
