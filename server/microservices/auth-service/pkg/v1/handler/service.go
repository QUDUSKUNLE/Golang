package handler

import (
	interfaces "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
	"google.golang.org/grpc"
)

type UserServiceStruct struct {
	useCase interfaces.UseCaseInterface
	user.UnimplementedUserServiceServer
	organization.UnimplementedOrganizationServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserServiceStruct{useCase: usecase}
	user.RegisterUserServiceServer(grpcServer, userGrpc)
	organization.RegisterOrganizationServiceServer(grpcServer, userGrpc)
}
