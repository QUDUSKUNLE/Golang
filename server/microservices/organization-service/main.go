package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/QUDUSKUNLE/microservices/organization-service/adapters/handler"
	"github.com/QUDUSKUNLE/microservices/organization-service/adapters/organizationcase"
	"github.com/QUDUSKUNLE/microservices/organization-service/adapters/subscribe"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/events"
	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
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
	// Load configuration
	// Load environment variable
	cfg, err := utils.LoadConfig()
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
		log.Fatalf("Error starting organization service: %v", err)
	}
	defer listen.Close()

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		middleware.ValidationInterceptor(),
	))
	go events.EventsSubscriber(cfg.KafkaBroker, cfg.KafkaTopic, cfg.KafkaGroup, subscribe.ProcessEvent)
	organizationUseCase := organizationcase.InitOrganizationServer(dbase)
	handler.NewOrganizationServer(grpcServer, organizationUseCase)
	go func() {
		logger.GetLogger().Info("Organization Service listening at with TLS enabled (Min version: TLS 1.2)", zap.String("address", cfg.Port))

		if err := grpcServer.Serve(listen); err != nil {
			logger.GetLogger().Fatal("failed to serve organization service", zap.Error(err))
		}
	}()
	<-ctx.Done()
	logger.GetLogger().Info("Shutting down organization service")
	grpcServer.GracefulStop()

	logger.GetLogger().Info("Organization service stopped")
}
