package handlers

import (
	"github.com/QUDUSKUNLE/microservices/shipping-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/shipping-service/protogen/golang/shipping"
	"google.golang.org/grpc"
)

type Handler struct {
	internalServicesAdapter ports.ShippingPorts
	shipping.UnimplementedShippingServiceServer
}

func NewServer(grpcServer *grpc.Server, handlers ports.ShippingPorts) {
	shippingGrpc := &Handler{internalServicesAdapter: handlers}
	shipping.RegisterShippingServiceServer(grpcServer, shippingGrpc)
}
