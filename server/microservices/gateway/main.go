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
	"google.golang.org/grpc/metadata"

	"github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/QUDUSKUNLE/microservices/gateway/config"
	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
	"github.com/QUDUSKUNLE/microservices/record-service/protogen/golang/record"
)

func main() {
	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Initialize runtime server
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs("authorization", req.Header.Get("Authorization"))
		}),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register AuthServiceHandler
	if err := user.RegisterUserServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("AUTH"), opts); err != nil {
		log.Fatalf("Failed to register the user service handler: %v", err)
	}

	// Register OrganizationServiceHandler
	if err := organization.RegisterOrganizationServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("ORGANIZATION"), opts); err != nil {
		log.Fatalf("Failed to register the organization service handler: %v", err)
	}

	// Register RecordServiceHandler
	if err := record.RegisterRecordServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("RECORD"), opts); err != nil {
		log.Fatalf("Failed to register the record service handler: %v", err)
	}

	addr := fmt.Sprintf("%v:%v", os.Getenv("GATEWAY"), os.Getenv("GATEWAY_PORT"))
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Gateway server closed abruptly: %v", err)
	}
}
