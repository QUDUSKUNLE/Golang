package handler

import (
	v1 "github.com/QUDUSKUNLE/microservices/auth-service/adapters"
	// "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/client"
	"github.com/QUDUSKUNLE/microservices/events-service/domain"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/auth"
	"google.golang.org/grpc"
)

type AuthServiceStruct struct {
	authService         v1.AuthPorts
	eventBroker         domain.EventPorts
	auth.UnimplementedAuthServiceServer
}

func NewAuthServer(server *grpc.Server, usecase v1.AuthPorts) {
	AuthServiceController := &AuthServiceStruct{
		authService: usecase,
	}
	auth.RegisterAuthServiceServer(server, AuthServiceController)
}
