package ports

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/gateway/db"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
)

type RepositoryPorts interface {
	CreateOrganization(ctx context.Context, user domain.OrganizationDto) (*db.Organization, error)
	GetOrganization(ctx context.Context, id string) (*db.Organization, error)
	GetOrganizationByUserID(ctx context.Context, user_id string) (*db.Organization, error)
}

type UseCasePorts interface {
	CreateOrganization(ctx context.Context, user domain.OrganizationDto) (*db.Organization, error)
	GetOrganization(ctx context.Context, id string) (*db.Organization, error)
	GetOrganizationByUserID(ctx context.Context, user_id string) (*db.Organization, error)
}
