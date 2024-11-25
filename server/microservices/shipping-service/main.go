package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/QUDUSKUNLE/microservices/shipping-service/core/services"
	config "github.com/QUDUSKUNLE/microservices/shipping-service/internal/config"
	"github.com/QUDUSKUNLE/microservices/shipping-service/internal/handlers"
	"github.com/QUDUSKUNLE/microservices/shipping-service/internal/repository"
)

var (
	internal *services.ServicesHandler
)

func init() {
	// Load environment variable
	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	port := os.Getenv("PORT")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error starting auth service: %v", err)
	}

	internalStore, err := repository.OpenDBConnection()
	if err != nil {
		log.Fatalf("Error connecting to the databse: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	shippingPorts := repository.InitUserServer(internalStore)
	handlers.NewServer(grpcServer, shippingPorts)

	log.Printf("Shipping Service listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve shipping service: %v", err)
	}
}
