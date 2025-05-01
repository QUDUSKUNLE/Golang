package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/record-service/adapters/handler"
	"github.com/QUDUSKUNLE/microservices/record-service/core/services"
	"github.com/QUDUSKUNLE/microservices/shared/db"
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
	dbase := db.DatabaseConnection(cfg.DB_URL)

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
	recordUseCase := services.InitializeRecordService(dbase)
	handler.NewRecordServer(grpcServer, recordUseCase, os.Getenv("ORGANIZATION"), os.Getenv("USER_SERVICE"))
	logger.GetLogger().Info("Record Service listening on with TLS enabled (Min version: TLS 1.2)", zap.String("address", cfg.Port))
	if err := grpcServer.Serve(listen); err != nil {
		logger.GetLogger().Fatal("failed to serve record service", zap.Error(err))
	}
}
