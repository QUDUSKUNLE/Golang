package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (handler *OrganizationServiceStruct) CreateOrganization(ctx context.Context, req *organization.CreateOrganizationRequest) (*organization.CreateOrganizationResponse, error) {
	data := handler.transformOrganizationRPC(req)
	_organization, err := handler.organizationService.CreateOrganization(ctx, data)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	return &organization.CreateOrganizationResponse{Id: _organization.ID, CreatedAt: _organization.CreatedAt.Time.String(), UpdatedAt: _organization.UpdatedAt.Time.String()}, nil
}

func (srv *OrganizationServiceStruct) GetOrganization(ctx context.Context, req *organization.GetOrganizationRequest) (*organization.GetOrganizationResponse, error) {
	_organization, err := srv.organizationService.GetOrganization(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Organization not found")
	}
	return &organization.GetOrganizationResponse{
		Id: _organization.ID, UserId: _organization.UserID, CreatedAt: _organization.CreatedAt.Time.String(), UpdatedAt: _organization.UpdatedAt.Time.String()}, nil
}

func (srv *OrganizationServiceStruct) GetOrganizationByUserID(ctx context.Context, req *organization.GetOrganizationByUserIDRequest) (*organization.GetOrganizationResponse, error) {
	_organization, err := srv.organizationService.GetOrganizationByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Organization not found")
	}
	return &organization.GetOrganizationResponse{
		Id: _organization.ID, UserId: _organization.UserID, CreatedAt: _organization.CreatedAt.Time.String(), UpdatedAt: _organization.UpdatedAt.Time.String()}, nil
}
