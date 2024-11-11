package grpc

import (
	interfaces "github.com/QUDUSKUNLE/microservices/pkg/v1"
	pb "github.com/QUDUSKUNLE/microservices/proto"
	"google.golang.org/grpc"
)

type UserServiceStruct struct {
	useCase interfaces.UseCaseInterface
	pb.UnimplementedUserServiceServer
	pb.UnimplementedGreetServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserServiceStruct{ useCase: usecase }
	pb.RegisterUserServiceServer(grpcServer, userGrpc)
	pb.RegisterGreetServiceServer(grpcServer, userGrpc)
}
