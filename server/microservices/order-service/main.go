package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QUDUSKUNLE/microservices/order-service/internal/config"
	dbconfig "github.com/QUDUSKUNLE/microservices/order-service/internal/db"
	handler "github.com/QUDUSKUNLE/microservices/order-service/pkg/v1/handler"
	"google.golang.org/grpc"
)

func main() {
	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}

	host, user, dbname, password, port := os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("PORT")
	db := dbconfig.DbConn(host, user, dbname, password)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error STARTING THE SERVER : %v", err)
	}

	grpcServer := grpc.NewServer()

	userUseCase := dbconfig.InitUserServer(db)
	handler.NewServer(grpcServer, userUseCase)

	log.Printf("Order Service listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
