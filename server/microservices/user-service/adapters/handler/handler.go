package handler

import (
	v1 "github.com/QUDUSKUNLE/microservices/user-service/adapters"

	"github.com/QUDUSKUNLE/microservices/events-service/domain"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"google.golang.org/grpc"
)

type UserServiceStruct struct {
	userService         v1.UserPorts
	eventBroker         domain.EventPorts
	user.UnimplementedUserServiceServer
}

func NewUserService(server *grpc.Server, usecase v1.UserPorts, brok domain.EventPorts) {
	userServiceController := &UserServiceStruct{
		userService:         usecase,
		eventBroker:         brok,
	}
	user.RegisterUserServiceServer(server, userServiceController)
}
