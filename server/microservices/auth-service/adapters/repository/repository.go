package repository

import (
	"context"

	v1 "github.com/QUDUSKUNLE/microservices/auth-service/adapters"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	database *db.Queries
}

// CreateUser implements v1.UseCaseInterface.
func (u *Repository) CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error) {
	return u.database.CreateUser(ctx, user)
}

// GetUser implements v1.UseCaseInterface.
func (u *Repository) GetUser(ctx context.Context, id string) (*db.User, error) {
	return u.database.GetUser(ctx, id)
}

func (u *Repository) GetUserByEmail(ctx context.Context, email pgtype.Text) (*db.User, error) {
	return u.database.GetUserByEmail(ctx, email)
}


func NewRepository(dbase *db.Queries) v1.AuthRepository {
	return &Repository{database: dbase}
}
