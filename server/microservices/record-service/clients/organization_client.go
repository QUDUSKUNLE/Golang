package clients

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/organization-service/db"
	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
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
