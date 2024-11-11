package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/QUDUSKUNLE/microservices/proto"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Greet(ctx, &pb.GreetRequest{Name: "Hello world"})
	if err != nil {
		log.Fatalf("Could not send greeting: %v", err)
	}

	log.Printf("Greeting: %s", response.Message)
}
