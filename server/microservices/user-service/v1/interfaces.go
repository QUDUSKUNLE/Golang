package v1

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/db"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error)
	GetUser(ctx context.Context, id string) (*db.User, error)
	GetUsers(ctx context.Context, params db.GetUsersParams) ([]*db.User, error)
	GetUserByEmail(ctx context.Context, email string) (*db.User, error)
	UpdateUser(ctx context.Context, user db.UpdateUserParams) (*db.UpdateUserRow, error)
}

type UserPorts interface {
	CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error)
	GetUser(ctx context.Context, id string) (*db.User, error)
	GetUsers(ctx context.Context, params db.GetUsersParams) ([]*db.User, error)
	GetUserByEmail(ctx context.Context, email string) (*db.User, error)
	UpdateUser(ctx context.Context, user db.UpdateUserParams) (*db.UpdateUserRow, error)
}
