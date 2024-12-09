package main

import (
	"fmt"
	"log"
	"net"
	"os"
	// "context"

	"github.com/QUDUSKUNLE/microservices/auth-service/internal/config"
	dbconfig "github.com/QUDUSKUNLE/microservices/auth-service/internal/db"
	handler "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/handler"
	"github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// "github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
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

	organization_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("ORGANIZATION_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer organization_conn.Close()

	grpcServer := grpc.NewServer()

	userUseCase := usecase.InitUserServer(db)
	handler.NewServer(grpcServer, userUseCase)

	log.Printf("Auth Service listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve auth service: %v", err)
	}
}
