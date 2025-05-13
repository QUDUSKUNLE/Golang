package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (srv *AuthServiceStruct) Signin(ctx context.Context, req *auth.SignInRequest) (*auth.SignInResponse, error) {
	user, err := srv.authService.Login(ctx, dto.LogInDto{Email: req.GetEmail(), Password: req.GetPassword()})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrInvalidCredentials)
	}

	if err = dto.ComparePassword(*user, req.GetPassword()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	token, err := srv.transformToken(dto.CurrentUser{
		ID:       user.ID,
		UserType: string(user.UserType),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token: %v", err)
	}
	return &auth.SignInResponse{Token: token}, nil
}
