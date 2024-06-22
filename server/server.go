package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
 	"google.golang.org/grpc/reflection"

	pb "tutorial/protocols"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := &pb.User{
		Id: req.Id,
		Name: "Abdul-Muhsin",
		Email: "aderemi.abdulmuhsin@gmail.com",
	}
	return &pb.UserResponse{User: user}, nil
}

func main() {
	list, err := net.Listen("tcp", ":50051");
	if err != nil {
		log.Fatalf("Fatal to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)

	log.Printf("Server is running on port 50051")
	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
