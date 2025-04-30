package v1

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthRepository interface {
	GetUser(ctx context.Context, id string) (*db.User, error)
	GetUserByEmail(ctx context.Context, email pgtype.Text) (*db.User, error)
}

type AuthPorts interface {
	GetUser(ctx context.Context, id string) (*db.User, error)
	Login(ctx context.Context, user dto.LogInDto) (*db.User, error)
}
