package repo

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/db"
	v1 "github.com/QUDUSKUNLE/microservices/user-service/v1"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	database *db.Queries
}

// GetUsers implements v1.RepositoryInterface.
func (u *Repository) GetUsers(ctx context.Context, params db.GetUsersParams) ([]*db.User, error) {
	return u.database.GetUsers(ctx, params)
}

// CreateUser implements v1.UseCaseInterface.
func (u *Repository) CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error) {
	return u.database.CreateUser(ctx, user)
}

// GetUser implements v1.UseCaseInterface.
func (u *Repository) GetUser(ctx context.Context, id string) (*db.User, error) {
	return u.database.GetUser(ctx, id)
}

func (u *Repository) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	emailText := pgtype.Text{String: email, Valid: true}
	return u.database.GetUserByEmail(ctx, emailText)
}

func (u *Repository) UpdateNin(ctx context.Context, user db.UpdateNinParams) (*db.UpdateNinRow, error) {
	return u.database.UpdateNin(ctx, user)
}

func NewRepository(dbase *db.Queries) v1.UserRepository {
	return &Repository{database: dbase}
}
