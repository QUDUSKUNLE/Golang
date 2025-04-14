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
	db := dbconfig.DbConn()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("Error starting auth service: %v", err)
	}
	userUseCase := usecase.InitUserServer(db)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// middleware.UnaryServerInterceptor(),
			middleware.ValidationInterceptor(),
		))
	handler.NewServer(grpcServer, userUseCase, os.Getenv("ORGANIZATION"))
	reflection.Register(grpcServer)

	log.Printf("Auth Service listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve auth service: %v", err)
	}
}
