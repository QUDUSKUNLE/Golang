package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/auth-service/internal/config"
	dbconfig "github.com/QUDUSKUNLE/microservices/auth-service/internal/db"
	handler "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/handler"
	"github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/usecase"
	"google.golang.org/grpc"
)

func init() {
	// Load environment variable
	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	host, user, dbname, password, port := os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("PORT")
	db := dbconfig.DbConn(host, user, dbname, password)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error starting auth service: %v", err)
	}

	grpcServer := grpc.NewServer()

	userUseCase := usecase.InitUserServer(db)
	handler.NewServer(grpcServer, userUseCase)

	log.Printf("Auth Service listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve auth service: %v", err)
	}
}
