package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/organization-service/adapters/handler"
	"github.com/QUDUSKUNLE/microservices/organization-service/adapters/organizationcase"
	"github.com/QUDUSKUNLE/microservices/organization-service/consumers"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
		"github.com/QUDUSKUNLE/microservices/shared/logger"
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
	dbase := db.DatabaseConnection()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("Error starting organization service: %v", err)
	}

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		middleware.ValidationInterceptor(),
	))
	go consumers.ConsumeCreatedUserEvent(dbase, os.Getenv("KAFKA_BROKER"), os.Getenv("KAFKA_TOPIC"), os.Getenv("KAFKA_GROUP_ID"))
	organizationUseCase := organizationcase.InitOrganizationServer(dbase)
	handler.NewOrganizationServer(grpcServer, organizationUseCase)
	logger.GetLogger().Info("Organization Service listening at with TLS enabled (Min version: TLS 1.2)", zap.Error(err))

	if err := grpcServer.Serve(listen); err != nil {
		logger.GetLogger().Fatal("failed to serve organization service", zap.Error(err))
	}
}
