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

	// Auth service
	auth_conn, err := grpc.NewClient(os.Getenv("AUTH"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer auth_conn.Close()

	// Organization service
	organization_conn, err := grpc.NewClient(os.Getenv("ORGANIZATION"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer organization_conn.Close()

	// Record service
	record_conn, err := grpc.NewClient(os.Getenv("RECORD"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer record_conn.Close()

	// Initialize runtime server
	mux := runtime.NewServeMux()

	// Register AuthServiceHandler
	if err = user.RegisterUserServiceHandler(context.Background(), mux, auth_conn); err != nil {
		log.Fatalf("Failed to register the user service handler: %v", err)
	}

	// Register OrganizationServiceHandler
	if err = organization.RegisterOrganizationServiceHandler(context.Background(), mux, organization_conn); err != nil {
		log.Fatalf("Failed to register the organizatin service handler: %v", err)
	}

	// Register RecordServiceHandler
	if err = record.RegisterRecordServiceHandler(context.Background(), mux, record_conn); err != nil {
		log.Fatalf("Failed to register the record service handler: %v", err)
	}

	addr := fmt.Sprintf("%v:%v", os.Getenv("GATEWAY"), os.Getenv("GATEWAY_PORT"))
	fmt.Println("Gateway server running on port: " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Gateway server closed abruptly: %v", err)
	}
}
