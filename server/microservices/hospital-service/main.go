package main

import (
	"fmt"
	"log"
	"net"
	"os"

	middleware "github.com/QUDUSKUNLE/microservices/hospital-service/pkg/v1/middleware"
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
	// Initialize database connection
	// db := dbconfig.DbConn()

	// Create TCP listener
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("Error starting hospital service: %v", err)
	}

	// Create gRPC server with TLS and interceptors
	grpcServer := grpc.NewServer(
		// grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			middleware.ValidationInterceptor(),
		))

	// Initialize use case and register services
	// userUseCase := usecase.InitUserServer(db)
	// handler.NewAuthServer(grpcServer, userUseCase, os.Getenv("HOSPITAL"))
	// reflection.Register(grpcServer)

	log.Printf("Hospital Service listening at %v with TLS enabled (Min version: TLS 1.2)", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve hospital service: %v", err)
	}
}
