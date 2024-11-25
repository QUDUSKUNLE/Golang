package handler

import (
	interfaces "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"google.golang.org/grpc"
)

type UserServiceStruct struct {
	useCase interfaces.UseCaseInterface
	user.UnimplementedUserServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserServiceStruct{useCase: usecase}
	user.RegisterUserServiceServer(grpcServer, userGrpc)
}
