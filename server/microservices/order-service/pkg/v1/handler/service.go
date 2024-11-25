package handler

import (
	interfaces "github.com/QUDUSKUNLE/microservices/order-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/order-service/protogen/golang/orders"
	"google.golang.org/grpc"
)

type UserServiceStruct struct {
	useCase interfaces.UseCaseInterface
	orders.UnimplementedOrderServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserServiceStruct{useCase: usecase}
	orders.RegisterOrderServiceServer(grpcServer, userGrpc)
}
