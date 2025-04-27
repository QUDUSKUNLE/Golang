package handler

import (
	v1 "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/client"
	"github.com/QUDUSKUNLE/microservices/events-service/domain"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"google.golang.org/grpc"
)

type AuthServiceStruct struct {
	authService         v1.AuthPorts
	organizationService ports.OrganizationPorts
	eventBroker         domain.EventPorts
	user.UnimplementedUserServiceServer
}

func NewAuthServer(server *grpc.Server, usecase v1.AuthPorts, conn string) {
	AuthServiceController := &AuthServiceStruct{
		authService:         usecase,
		organizationService: client.NewGRPCOrganizationService(conn),
	}
	user.RegisterUserServiceServer(server, AuthServiceController)
}
