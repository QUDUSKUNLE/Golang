package grpc

import (
	interfaces "github.com/QUDUSKUNLE/microservices/services/auth-service/pkg/v1"
	user "github.com/QUDUSKUNLE/microservices/services/auth-service/protogen/golang/user"
	greet "github.com/QUDUSKUNLE/microservices/services/auth-service/protogen/golang/greet"
	"google.golang.org/grpc"
)

type UserServiceStruct struct {
	useCase interfaces.UseCaseInterface
	user.UnimplementedUserServiceServer
	greet.UnimplementedGreetServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserServiceStruct{useCase: usecase}
	user.RegisterUserServiceServer(grpcServer, userGrpc)
	greet.RegisterGreetServiceServer(grpcServer, userGrpc)
}
