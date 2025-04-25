package clients

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/auth-service/adapters/dto"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
)

// CreateUser implements v1.UseCaseInterface.
func (u *userService) CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error) {
	panic("unimplemented")
}

// GetUser implements v1.UseCaseInterface.
func (this *userService) GetUser(ctx context.Context, id string) (*db.User, error) {
	response, err := this.userGrpcClient.Read(ctx, &user.SingleUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &db.User{ID: response.GetData().Id}, nil
}

// Login implements v1.UseCaseInterface.
func (u *userService) Login(ctx context.Context, user dto.LogInDto) (*db.User, error) {
	panic("unimplemented")
}

// UpdateNin implements v1.UseCaseInterface.
func (u *userService) UpdateNin(ctx context.Context, user db.UpdateNinParams) (*db.UpdateNinRow, error) {
	panic("unimplemented")
}
