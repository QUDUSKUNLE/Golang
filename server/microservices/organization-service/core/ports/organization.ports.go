package ports

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/organization-service/db"
)

type RepositoryPorts interface {
	CreateOrganization(ctx context.Context, user domain.OrganizationDto) (*db.Organization, error)
	GetOrganization(ctx context.Context, id string) (*db.Organization, error)
}

type UseCasePorts interface {
	CreateOrganization(ctx context.Context, user domain.OrganizationDto) (*db.Organization, error)
	GetOrganization(ctx context.Context, id string) (*db.Organization, error)
}
