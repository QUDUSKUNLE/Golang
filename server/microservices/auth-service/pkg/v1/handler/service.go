package handler

import (
	interfaces "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/client"
	"github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"google.golang.org/grpc"
)

type UserServiceStruct struct {
	userService         interfaces.UseCaseInterface
	organizationService ports.UseCasePorts
	user.UnimplementedUserServiceServer
}

func NewServer(server *grpc.Server, usecase interfaces.UseCaseInterface, conn string) {
	userServiceController := &UserServiceStruct{
		userService:         usecase,
		organizationService: client.NewGRPCOrganizationService(conn),
	}
	user.RegisterUserServiceServer(server, userServiceController)
}
