package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/organization-service/adapters/handler"
	"github.com/QUDUSKUNLE/microservices/organization-service/adapters/config"
	"github.com/QUDUSKUNLE/microservices/organization-service/adapters/usecase"
	"github.com/QUDUSKUNLE/microservices/organization-service/db"
	"google.golang.org/grpc"
)

func init() {
	// Load environment variable
	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbase := db.DatabaseConnection()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("Error starting organization service: %v", err)
	}

	grpcServer := grpc.NewServer()
	organizationUseCase := usecase.InitOrganizationServer(dbase)
	handler.NewOrganizationServer(grpcServer, organizationUseCase)
	log.Printf("Organization Service listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve organization service: %v", err)
	}
}
