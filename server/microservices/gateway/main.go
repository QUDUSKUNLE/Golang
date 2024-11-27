package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/QUDUSKUNLE/microservices/gateway/config"
)

var request_count = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "gateway",
		Name:      "http_request_count",
		Help:      "Tracks the total number of HTTP requests received",
	},
)

func main() {
	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/prometheus", func(w http.ResponseWriter, r *http.Request) {
		request_count.Inc()
	})

	auth_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("AUTH_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer auth_conn.Close()

	organization_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("ORGANIZATION_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer organization_conn.Close()

	shipping_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("SHIPPING_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer shipping_conn.Close()

	mux := runtime.NewServeMux()

	if err = user.RegisterUserServiceHandler(context.Background(), mux, auth_conn); err != nil {
		log.Fatalf("Failed to register the user service handler: %v", err)
	}
	addr := fmt.Sprintf("%v:%v", os.Getenv("GATEWAY"), os.Getenv("GATEWAY_PORT"))
	fmt.Println("Gateway server running on port: " + addr)
	prometheus.MustRegister(request_count)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Gateway server closed abruptly: %v", err)
	}
}
