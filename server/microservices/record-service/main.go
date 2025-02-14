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
	"github.com/QUDUSKUNLE/microservices/record-service/db"
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

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(middleware.StreamInterceptor()),
		grpc.ChainUnaryInterceptor(
			middleware.UnaryServerInterceptor(),
			middleware.ValidateUnaryInterceptor(),
		),
	)
	recordUseCase := usecase.InitializeRecordService(dbase)
	handler.NewRecordServer(grpcServer, recordUseCase, os.Getenv("ORGANIZATION"), os.Getenv("AUTH"))
	log.Printf("Record Service listening on %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve record service: %v", err)
	}
}
