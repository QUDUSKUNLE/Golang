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
	_organization, err := handler.organizationService.CreateOrganization(ctx, data)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	return &organization.CreateOrganizationResponse{Id: _organization.ID, CreatedAt: _organization.CreatedAt.Time.String(), UpdatedAt: _organization.UpdatedAt.Time.String()}, nil
}

func (srv *OrganizationServiceStruct) GetOrganization(ctx context.Context, req *organization.GetOrganizationRequest) (*organization.GetOrganizationResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	_organization, err := srv.organizationService.GetOrganization(ctx, id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Organization not found")
	}
	return &organization.GetOrganizationResponse{
		Id: _organization.ID, UserId: _organization.UserID, CreatedAt: _organization.CreatedAt.Time.String(), UpdatedAt: _organization.UpdatedAt.Time.String()}, nil
}
