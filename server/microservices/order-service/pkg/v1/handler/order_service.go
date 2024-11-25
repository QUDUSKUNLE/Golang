package handler

import (
	"context"
	"log"

	"github.com/QUDUSKUNLE/microservices/order-service/protogen/golang/orders"
)

func (srv *UserServiceStruct) AddOrder(ctx context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received an add-order request")
	return &orders.Empty{}, nil
}
