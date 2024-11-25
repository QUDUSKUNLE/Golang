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
	"github.com/QUDUSKUNLE/microservices/order-service/protogen/golang/orders"
	"github.com/QUDUSKUNLE/microservices/shipping-service/protogen/golang/shipping"
)

func main() {
	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}

	auth_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("AUTH_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer auth_conn.Close()

	order_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("ORDER_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer order_conn.Close()

	shipping_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("SHIPPING_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer shipping_conn.Close()

	mux := runtime.NewServeMux()

	if err = user.RegisterUserServiceHandler(context.Background(), mux, auth_conn); err != nil {
		log.Fatalf("Failed to register the user service handler: %v", err)
	}

	if err := orders.RegisterOrderServiceHandler(context.Background(), mux, order_conn); err != nil {
		log.Fatalf("Failed to register the order service handler: %v", err)
	}

	if err := shipping.RegisterShippingServiceHandler(context.Background(), mux, order_conn); err != nil {
		log.Fatalf("Failed to register the shipping service handler: %v", err)
	}

	addr := fmt.Sprintf("%v:%v", os.Getenv("GATEWAY"), os.Getenv("GATEWAY_PORT"))
	fmt.Println("Gateway server running on port: " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Gateway server closed abruptly: %v", err)
	}
}
