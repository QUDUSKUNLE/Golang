package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/auth-service/internal/db"
	"github.com/QUDUSKUNLE/microservices/auth-service/internal/dto"
	userProtoc "github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
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

	built_user, err := dto.BuildNewUser(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	user, err := srv.useCase.CreateUser(
		ctx, db.CreateUserParams{
			Email:    built_user.Email,
			Nin:      pgtype.Text{String: "", Valid: true},
			Password: built_user.Password,
			UserType: built_user.UserType,
		})
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	if data.UserType == db.UserEnumORGANIZATION {
		_, err := srv.organizationService.CreateOrganization(ctx, domain.OrganizationDto{UserID: user.ID})
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
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
		return nil, status.Error(codes.NotFound, "incorrect log in credentials")
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

func (srv *UserServiceStruct) UpdateNin(ctx context.Context, req *userProtoc.UpdateNinRequest) (*userProtoc.UpdateNinResponse, error) {
	UserID, ok := ctx.Value("UserID").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized to perform operation.")
	}
	nin := req.GetNin()
	if nin == "" {
		return nil, status.Error(codes.InvalidArgument, Nin_Required)
	}
	_, err := srv.useCase.UpdateNin(ctx, db.UpdateNinParams{
		Nin: pgtype.Text{
			String: nin, Valid: true,
		},
		ID: UserID})
	if err != nil {
		return nil, status.Error(codes.Unimplemented, err.Error())
	}
	return &userProtoc.UpdateNinResponse{Data: "Nin updated successfully"}, nil
}

func (srv *UserServiceStruct) Home(ctx context.Context, req *userProtoc.HomeRequest) (*userProtoc.GetHomeResponse, error) {
	return &userProtoc.GetHomeResponse{Message: Welcome_Home}, nil
}
