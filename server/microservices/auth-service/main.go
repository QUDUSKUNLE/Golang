package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/auth-service/adapters/config"
	dbconfig "github.com/QUDUSKUNLE/microservices/auth-service/adapters/db"
	handler "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/handler"
	middleware "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/middleware"
	"github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	// Load environment variable
	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Initialize database connection
	db := dbconfig.DbConn()

	// Create TCP listener
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("Error starting auth service: %v", err)
	}

	// Create gRPC server with TLS and interceptors
	grpcServer := grpc.NewServer(
		// grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			middleware.ValidationInterceptor(),
		))

	// Initialize use case and register services
	userUseCase := usecase.InitUserServer(db)
	handler.NewAuthServer(grpcServer, userUseCase, os.Getenv("ORGANIZATION"))
	reflection.Register(grpcServer)

	log.Printf("Auth Service listening at %v with TLS enabled (Min version: TLS 1.2)", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve auth service: %v", err)
	}
}
