package handler

import (
	"context"
	"errors"

	"github.com/QUDUSKUNLE/microservices/auth-service/internal/db"
	"github.com/QUDUSKUNLE/microservices/auth-service/internal/dto"
	userProtoc "github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/jackc/pgx/v5/pgtype"
)

func (srv *UserServiceStruct) Create(ctx context.Context, req *userProtoc.CreateUserRequest) (*userProtoc.SuccessResponse, error) {
	data := srv.transformUserRPC(req)
	if data.Email == "" || data.Password == "" || string(data.UserType) == "" {
		return &userProtoc.SuccessResponse{}, errors.New("please provide all fields")
	} else if data.Password != data.ConfirmPassword {
		return &userProtoc.SuccessResponse{}, errors.New("incorrect passwords")
	}
	user, err := dto.BuildNewUser(data)
	if err != nil {
		return &userProtoc.SuccessResponse{}, err
	}
	_, err = srv.useCase.CreateUser(
		ctx, db.CreateUserParams{
			Email:    user.Email,
			Nin:      pgtype.Text{String: "", Valid: true},
			Password: user.Password,
			UserType: user.UserType,
		})
	if err != nil {
		return &userProtoc.SuccessResponse{}, err
	}
	return &userProtoc.SuccessResponse{Data: "User registered successfully."}, nil
}

func (srv *UserServiceStruct) Read(ctx context.Context, req *userProtoc.SingleUserRequest) (*userProtoc.GetUserResponse, error) {
	id := req.GetId()
	if id == "" {
		return &userProtoc.GetUserResponse{}, errors.New("id cannot be blank")
	}
	user, err := srv.useCase.GetUser(ctx, id)
	if err != nil {
		return &userProtoc.GetUserResponse{}, errors.New("user not found")
	}
	data := &userProtoc.User{Id: user.ID, Email: user.Email.String, CreatedAt: user.CreatedAt.Time.String(), UpdatedAt: user.UpdatedAt.Time.String()}
	return &userProtoc.GetUserResponse{
		Data: data}, nil
}

func (srv *UserServiceStruct) Signin(ctx context.Context, req *userProtoc.SignInRequest) (*userProtoc.SignInResponse, error) {
	email, password := req.GetEmail(), req.GetPassword()

	if email == "" || password == "" {
		return &userProtoc.SignInResponse{}, errors.New("please provide all fields")
	}
	user, err := srv.useCase.Login(ctx, dto.LogInDto{Email: email, Password: password})
	if err != nil {
		return &userProtoc.SignInResponse{}, err
	}

	if err = dto.ComparePassword(*user, password); err != nil {
		return &userProtoc.SignInResponse{}, err
	}
	token, err := srv.transformToken(dto.CurrentUser{
		ID:       user.ID,
		UserType: db.NullUserEnum{UserEnum: user.UserType},
	})
	return &userProtoc.SignInResponse{Token: token}, nil
}

func (srv *UserServiceStruct) Home(ctx context.Context, req *userProtoc.HomeRequest) (*userProtoc.GetHomeResponse, error) {
	return &userProtoc.GetHomeResponse{Message: "Welcome to Bahsoon Shipping Inc."}, nil
}
