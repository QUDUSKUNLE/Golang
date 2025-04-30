package services

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/db"
	repo "github.com/QUDUSKUNLE/microservices/user-service/core/repository"
	v1 "github.com/QUDUSKUNLE/microservices/user-service/v1"
)

type UserCase struct {
	database v1.UserRepository
}

// GetUsers implements v1.UseCaseInterface.
func (u *UserCase) GetUsers(ctx context.Context, params db.GetUsersParams) ([]*db.User, error) {
	return u.database.GetUsers(ctx, params)
}

// UpdateNin implements v1.UseCaseInterface.
func (u *UserCase) UpdateNin(ctx context.Context, data db.UpdateNinParams) (*db.UpdateNinRow, error) {
	return u.database.UpdateNin(ctx, data)
}

// GetUserByEmail implements v1.RepositoryInterface.
func (u *UserCase) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	return u.database.GetUserByEmail(ctx, email)
}

// CreateUser implements v1.UseCaseInterface.
func (u *UserCase) CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error) {
	return u.database.CreateUser(ctx, user)
}

// GetUser implements v1.UseCaseInterface.
func (u *UserCase) GetUser(ctx context.Context, id string) (*db.User, error) {
	return u.database.GetUser(ctx, id)
}

func New(repo v1.UserRepository) v1.UserRepository {
	return &UserCase{database: repo}
}

func InitUserServer(db *db.Queries) v1.UserPorts {
	userRepo := repo.NewRepository(db)
	return &UserCase{database: userRepo}
}
