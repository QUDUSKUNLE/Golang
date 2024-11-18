package handler

import (
	"context"
	"errors"

	"github.com/QUDUSKUNLE/microservices/services/auth-service/internal/models"
	userProtoc "github.com/QUDUSKUNLE/microservices/services/auth-service/protogen/golang/user"
)

func (srv *UserServiceStruct) Create(ctx context.Context, req *userProtoc.CreateUserRequest) (*userProtoc.SuccessResponse, error) {
	data := srv.transformUserRPC(req)
	if data.Email == "" || data.Password == "" {
		return &userProtoc.SuccessResponse{}, errors.New("please provide all fields")
	} else if data.Password != req.GetConfirmPassword() {
		return &userProtoc.SuccessResponse{}, errors.New("incorrect passwords")
	}

	if err := srv.useCase.CreateUser(data); err != nil {
		return &userProtoc.SuccessResponse{}, err
	}
	return  &userProtoc.SuccessResponse{Data: "User registered successfully."}, nil
}

func (srv *UserServiceStruct) ReadUsers(ctx context.Context, req *userProtoc.GetUsersRequest) (*userProtoc.GetUsersResponse, error) {
	users, err := srv.useCase.GetUsers()
	if err != nil {
		return &userProtoc.GetUsersResponse{}, errors.New("users not found")
	}
	return srv.transformUsers(users), nil
}

func (srv *UserServiceStruct) Read(ctx context.Context, req *userProtoc.SingleUserRequest) (*userProtoc.GetUserResponse, error) {
	id := req.GetId()
	if id == "" {
		return &userProtoc.GetUserResponse{}, errors.New("id cannot be blank")
	}

	user, err := srv.useCase.GetUser(id)
	if err != nil {
		return &userProtoc.GetUserResponse{}, errors.New("user not found")
	}
	data := &userProtoc.User{Id: user.ID.String(), Email: user.Email, CreatedAt: user.CreatedAt.String(), UpdatedAt: user.UpdatedAt.Local().String()}
	return &userProtoc.GetUserResponse{
		Data: data }, nil
}

func (srv *UserServiceStruct) Signin(ctx context.Context, req *userProtoc.SignInRequest) (*userProtoc.SignInResponse, error) {
	email, password := req.GetEmail(), req.GetPassword();

	if email == "" || password == "" {
		return &userProtoc.SignInResponse{}, errors.New("please provide all fields")
	}
	user, err := srv.useCase.LogIn(models.LogInDto{Email: email, Password: password})
	if err != nil {
		return &userProtoc.SignInResponse{}, err
	}
	token, err := srv.transformToken(models.CurrentUser{ID: user.ID, UserType: models.USER.ReturnUserString()})
	return &userProtoc.SignInResponse{Token: token}, nil
}
