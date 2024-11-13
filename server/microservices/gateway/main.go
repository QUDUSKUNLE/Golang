package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/QUDUSKUNLE/microservices/services/order-service/protogen/golang/orders"
	"github.com/QUDUSKUNLE/microservices/services/auth-service/protogen/golang/user"
)

func main() {
	auth_conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer auth_conn.Close()

	order_conn, err := grpc.NewClient("127.0.0.1:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer order_conn.Close()

	mux := runtime.NewServeMux()

	if err = user.RegisterUserServiceHandler(context.Background(), mux, auth_conn); err != nil {
		log.Fatalf("Failed to register the user service handler: %v", err)
	}

	if err := orders.RegisterOrderServiceHandler(context.Background(), mux, order_conn); err != nil {
		log.Fatalf("Failed to register the order service handler: %v", err)
	}

	addr := "0.0.0.0:7556"
	fmt.Println("Gateway server running on port: " +  addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Gateway server closed abruptly: %v", err)
	}
}
