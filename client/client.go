package main

import (
	"context"
	"log"
	"time"

	pb "github.com/QUDUSKUNLE/Golang/tutorial/protocols"
	"google.golang.org/grpc"
)

func main() {
 conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
 if err != nil {
  log.Fatalf("Failed to connect: %v", err)
 }
 defer conn.Close()

 client := pb.NewUserServiceClient(conn)

 ctx, cancel := context.WithTimeout(context.Background(), time.Second)
 defer cancel()

 req := &pb.UserRequest{Id: 1}
 res, err := client.GetUser(ctx, req)
 if err != nil {
  log.Fatalf("Failed to get user: %v", err)
 }

 log.Printf("User: %v", res.User)
}
