package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/auth-service/internal/db"
	"github.com/QUDUSKUNLE/microservices/auth-service/internal/dto"
	userProtoc "github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (srv *UserServiceStruct) Create(ctx context.Context, req *userProtoc.CreateUserRequest) (*userProtoc.SuccessResponse, error) {
	data := srv.transformUserRPC(req)
	if data.Email == "" || data.Password == "" || string(data.UserType) == "" {
		return nil, status.Error(codes.InvalidArgument, All_Fields)
	} else if data.Password != data.ConfirmPassword {
		return nil, status.Error(codes.InvalidArgument, Incorrect_Password)
	}
	user, err := dto.BuildNewUser(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	_, err = srv.useCase.CreateUser(
		ctx, db.CreateUserParams{
			Email:    user.Email,
			Nin:      pgtype.Text{String: "", Valid: true},
			Password: user.Password,
			UserType: user.UserType,
		})
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	return &userProtoc.SuccessResponse{Data: Registered_Successfully}, nil
}

func (srv *UserServiceStruct) Read(ctx context.Context, req *userProtoc.SingleUserRequest) (*userProtoc.GetUserResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, Provide_ID)
	}
	user, err := srv.useCase.GetUser(ctx, id)
	if err != nil {
		return nil, status.Error(codes.NotFound, Not_Found)
	}
	data := &userProtoc.User{Id: user.ID, Email: user.Email.String, CreatedAt: user.CreatedAt.Time.String(), UpdatedAt: user.UpdatedAt.Time.String()}
	return &userProtoc.GetUserResponse{
		Data: data}, nil
}

func (srv *UserServiceStruct) Signin(ctx context.Context, req *userProtoc.SignInRequest) (*userProtoc.SignInResponse, error) {
	email, password := req.GetEmail(), req.GetPassword()

	if email == "" || password == "" {
		return nil, status.Error(codes.InvalidArgument, All_Fields)
	}
	user, err := srv.useCase.Login(ctx, dto.LogInDto{Email: email, Password: password})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err = dto.ComparePassword(*user, password); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	token, err := srv.transformToken(dto.CurrentUser{
		ID:       user.ID,
		UserType: db.NullUserEnum{UserEnum: user.UserType},
	})
	return &userProtoc.SignInResponse{Token: token}, nil
}

func (srv *UserServiceStruct) Home(ctx context.Context, req *userProtoc.HomeRequest) (*userProtoc.GetHomeResponse, error) {
	return &userProtoc.GetHomeResponse{Message: Welcome_Home}, nil
}
