package handler

import (
	v1 "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/client"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"google.golang.org/grpc"
)

type UserServiceStruct struct {
	userService         v1.UserPorts
	organizationService ports.OrganizationPorts
	user.UnimplementedUserServiceServer
}

func NewAuthServer(server *grpc.Server, usecase v1.UserPorts, conn string) {
	userServiceController := &UserServiceStruct{
		userService:         usecase,
		organizationService: client.NewGRPCOrganizationService(conn),
	}
	user.RegisterUserServiceServer(server, userServiceController)
}
