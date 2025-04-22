package usecase

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/auth-service/adapters/dto"
	interfaces "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	repo "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/repository"
	"github.com/QUDUSKUNLE/microservices/gateway/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type UseCase struct {
	repo interfaces.RepositoryInterface
}

// GetUsers implements v1.UseCaseInterface.
func (u *UseCase) GetUsers(ctx context.Context) ([]*db.User, error) {
	return u.repo.GetUsers(ctx)
}

// UpdateNin implements v1.UseCaseInterface.
func (u *UseCase) UpdateNin(ctx context.Context, data db.UpdateNinParams) (*db.UpdateNinRow, error) {
	return u.repo.UpdateNin(ctx, data)
}

// GetUserByEmail implements v1.RepositoryInterface.
func (u *UseCase) GetUserByEmail(ctx context.Context, email pgtype.Text) (*db.User, error) {
	panic("unimplemented")
}

// Login implements v1.UseCaseInterface.
func (u *UseCase) Login(ctx context.Context, user dto.LogInDto) (*db.User, error) {
	return u.repo.GetUserByEmail(ctx, pgtype.Text{String: user.Email, Valid: true})
}

// CreateUser implements v1.UseCaseInterface.
func (u *UseCase) CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error) {
	return u.repo.CreateUser(ctx, user)
}

// GetUser implements v1.UseCaseInterface.
func (u *UseCase) GetUser(ctx context.Context, id string) (*db.User, error) {
	return u.repo.GetUser(ctx, id)
}

func New(repo interfaces.RepositoryInterface) interfaces.RepositoryInterface {
	return &UseCase{repo: repo}
}

func InitUserServer(db *db.Queries) interfaces.UseCaseInterface {
	userRepo := repo.NewRepository(db)
	return &UseCase{repo: userRepo}
}
