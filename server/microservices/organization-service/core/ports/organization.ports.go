package ports

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
)

type RepositoryPorts interface {
	CreateOrganization(ctx context.Context, user dto.OrganizationDto) (*db.Organization, error)
	GetOrganization(ctx context.Context, id string) (*db.Organization, error)
	GetOrganizationByUserID(ctx context.Context, user_id string) (*db.Organization, error)
}

type OrganizationPorts interface {
	CreateOrganization(ctx context.Context, user dto.OrganizationDto) (*db.Organization, error)
	GetOrganization(ctx context.Context, id string) (*db.Organization, error)
	GetOrganizationByUserID(ctx context.Context, user_id string) (*db.Organization, error)
}
