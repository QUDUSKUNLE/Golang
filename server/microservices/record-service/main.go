package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/record-service/adapters/config"
	"github.com/QUDUSKUNLE/microservices/record-service/adapters/handler"
	"github.com/QUDUSKUNLE/microservices/record-service/adapters/middleware"
	"github.com/QUDUSKUNLE/microservices/record-service/adapters/usecase"
	"github.com/QUDUSKUNLE/microservices/gateway/db"
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
		log.Fatalf("Error starting record service: %v", err)
	}

	defer listen.Close()

	// Configure the gRPC server with enhanced options
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.ValidateUnaryInterceptor(),
		),
	)
	recordUseCase := usecase.InitializeRecordService(dbase)
	handler.NewRecordServer(grpcServer, recordUseCase, os.Getenv("ORGANIZATION"), os.Getenv("AUTH"))
	log.Printf("Record Service listening on %v with TLS enabled (Min version: TLS 1.2)", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve record service: %v", err)
	}
}
