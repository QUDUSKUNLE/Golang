package client

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/gateway/db"
	"github.com/QUDUSKUNLE/microservices/gateway/protogen/organization"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
)

// AuthInterceptor is a struct to hold the token for authorization
type AuthInterceptor struct {
	authToken string
}

func (o *organizationService) CreateOrganization(ctx context.Context, user domain.OrganizationDto) (*db.Organization, error) {
	req := &organization.CreateOrganizationRequest{UserId: user.UserID}
	resp, err := o.grpcClient.CreateOrganization(ctx, req)
	if err != nil {
		return &db.Organization{}, err
	}
	return &db.Organization{ID: resp.Id}, nil
}

// GetOrganization implements ports.UseCasePorts.
func (o *organizationService) GetOrganization(ctx context.Context, id string) (*db.Organization, error) {
	panic("unimplemented")
}

// GetOrganizationID implements ports.UseCasePorts.
func (o *organizationService) GetOrganizationByUserID(ctx context.Context, user_id string) (*db.Organization, error) {
	panic("unimplemented")
}
