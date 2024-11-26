package handlers

import (
	"context"

	shippingProtoc "github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/shipping"
)

func (handler *Handler) PostShipping(context context.Context, req *shippingProtoc.CreateShippingRequest) error {
	return nil
}
