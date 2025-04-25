package handler

import (
	v1 "github.com/QUDUSKUNLE/microservices/hospital-service/pkg/v1"
	// "github.com/QUDUSKUNLE/microservices/hospital-service/pkg/v1/client"
	// "github.com/QUDUSKUNLE/microservices/hospital-service/protogen/golang/user"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"google.golang.org/grpc"
)

type UserServiceStruct struct {
	userService         v1.UseCaseInterface
	organizationService ports.OrganizationPorts
	// user.UnimplementedUserServiceServer
}

func NewAuthServer(server *grpc.Server, usecase v1.UseCaseInterface, conn string) {
	// userServiceController := &UserServiceStruct{
	// 	userService:         usecase,
	// 	organizationService: client.NewGRPCOrganizationService(conn),
	// }
	// user.RegisterUserServiceServer(server, userServiceController)
}
