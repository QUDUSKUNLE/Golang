package handler

import (
	"github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/client"
	v1 "github.com/QUDUSKUNLE/microservices/user-service/v1"

	"github.com/QUDUSKUNLE/microservices/events-service/domain"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"google.golang.org/grpc"
)

type UserServiceStruct struct {
	userService         v1.UserPorts
	organizationService ports.OrganizationPorts
	eventBroker         domain.EventPorts
	user.UnimplementedUserServiceServer
}

func NewUserService(server *grpc.Server, usecase v1.UserPorts, brok domain.EventPorts, conn string) {
	userServiceController := &UserServiceStruct{
		userService:         usecase,
		organizationService: client.NewGRPCOrganizationService(conn),
		eventBroker:         brok,
	}
	user.RegisterUserServiceServer(server, userServiceController)
}
