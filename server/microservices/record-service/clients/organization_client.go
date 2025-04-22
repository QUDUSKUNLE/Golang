package clients

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/gateway/db"
	"github.com/QUDUSKUNLE/microservices/gateway/protogen/organization"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
)

// CreateOrganization implements ports.UseCasePorts.
func (this *organizationService) CreateOrganization(ctx context.Context, user domain.OrganizationDto) (*db.Organization, error) {
	req := &organization.CreateOrganizationRequest{UserId: user.UserID}
	resp, err := this.organizationGrpcClient.CreateOrganization(ctx, req)
	if err != nil {
		return &db.Organization{}, nil
	}
	return &db.Organization{ID: resp.Id}, nil
}

// GetOrganization implements ports.UseCasePorts.
func (this *organizationService) GetOrganization(ctx context.Context, id string) (*db.Organization, error) {
	panic("unimplemented")
}

// GetOrganizationID implements ports.UseCasePorts.
func (this *organizationService) GetOrganizationByUserID(ctx context.Context, user_id string) (*db.Organization, error) {
	req := &organization.GetOrganizationByUserIDRequest{UserId: user_id}
	res, err := this.organizationGrpcClient.GetOrganizationByUserID(ctx, req)
	if err != nil {
		return &db.Organization{}, nil
	}
	return &db.Organization{ID: res.Id, UserID: res.UserId}, nil
}
