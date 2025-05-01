package main

import (
	"fmt"
	"log"
	"net"
	"os"

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
	dbase := db.DatabaseConnection()
	port := os.Getenv("PORT")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error starting organization service: %v", err)
	}

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		middleware.ValidationInterceptor(),
	))
	go events.EventsSubscriber(os.Getenv("KAFKA_BROKER"), os.Getenv("KAFKA_TOPIC"), os.Getenv("KAFKA_GROUP_ID"), subscribe.ProcessEvent)
	organizationUseCase := organizationcase.InitOrganizationServer(dbase)
	handler.NewOrganizationServer(grpcServer, organizationUseCase)
	logger.GetLogger().Info("Organization Service listening at with TLS enabled (Min version: TLS 1.2)", zap.String("address", port))

	if err := grpcServer.Serve(listen); err != nil {
		logger.GetLogger().Fatal("failed to serve organization service", zap.Error(err))
	}
}
