package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/QUDUSKUNLE/microservices/gateway/config"
	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
	"github.com/QUDUSKUNLE/microservices/record-service/protogen/golang/record"
)

func main() {
	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize runtime server
	mux := runtime.NewServeMux()

	// Register AuthServiceHandler
	if err := user.RegisterUserServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		os.Getenv("AUTH"), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	); err != nil {
		log.Fatalf("Failed to register the user service handler: %v", err)
	}

	// Register OrganizationServiceHandler
	if err := organization.RegisterOrganizationServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		os.Getenv("ORGANIZATION"), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		log.Fatalf("Failed to register the organization service handler: %v", err)
	}

	// Register RecordServiceHandlerFromEndpoint
	if err := record.RegisterRecordServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		os.Getenv("RECORD"), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		log.Fatalf("Failed to register the record service handler: %v", err)
	}

	addr := fmt.Sprintf("%v:%v", os.Getenv("GATEWAY"), os.Getenv("GATEWAY_PORT"))
	fmt.Println("Gateway server running on port: " + addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Gateway server closed abruptly: %v", err)
	}
}
