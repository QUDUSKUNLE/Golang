package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/schedule-service/adapters/repository"
	"github.com/QUDUSKUNLE/microservices/schedule-service/adapters/server"

	"os/signal"
	"syscall"

	"github.com/QUDUSKUNLE/microservices/schedule-service/adapters/subscribe"
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
	cfg, err := utils.LoadConfig("SCHEDULE")
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
	repo := repository.NewScheduleRepository(dbase)
	server.NewScheduleServer(grpcServer, repo)

	// Run the gRPC server in a goroutine
	go func() {
		logger.GetLogger().Info("Schedule Service listening on with TLS enabled (Min version: TLS 1.2)", zap.String("address", cfg.Port))
		if err := grpcServer.Serve(listen); err != nil {
			logger.GetLogger().Fatal("failed to serve record service", zap.Error(err))
		}
	}()
	<-ctx.Done()
	logger.GetLogger().Info("Shutting down server gracefully...")
	grpcServer.GracefulStop()

	logger.GetLogger().Info("Schedule service stopped gracefully")
}
