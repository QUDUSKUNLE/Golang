package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (handler *OrganizationServiceStruct) CreateOrganization(ctx context.Context, req *organization.CreateOrganizationRequest) (*organization.CreateOrganizationResponse, error) {
	data := handler.transformOrganizationRPC(req)
	if data.UserID == "" {
		return nil, status.Error(codes.InvalidArgument, "User ID cannot be empty")
	}
	res, err := handler.useCase.CreateOrganization(ctx, data)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	return &organization.CreateOrganizationResponse{Id: res.ID, CreatedAt: res.CreatedAt.Time.String(), UpdatedAt: res.UpdatedAt.Time.String()}, nil
}
