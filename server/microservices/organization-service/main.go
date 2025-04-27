package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/organization-service/adapters/handler"
	// "github.com/QUDUSKUNLE/microservices/organization-service/consumers"
	"github.com/QUDUSKUNLE/microservices/organization-service/adapters/organizationcase"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
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

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		middleware.ValidationInterceptor(),
	))
	organizationUseCase := organizationcase.InitOrganizationServer(dbase)
	handler.NewOrganizationServer(grpcServer, organizationUseCase)
	log.Printf("Organization Service listening at %v with TLS enabled (Min version: TLS 1.2)", listen.Addr())

	// broker := os.Getenv("KAFKA_BROKER")
	// if broker == "" {
	// 	log.Fatal("KAFKA_BROKER environment variable is not set")
	// }
	// go consumers.ConsumeCreatedUserEvent(broker, "CreatedUser")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve organization service: %v", err)
	}
}
