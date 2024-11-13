package handler

import (
	"context"
	"errors"
	"fmt"

	u "github.com/QUDUSKUNLE/microservices/services/auth-service/protogen/golang/user"
)

func (srv *UserServiceStruct) Create(ctx context.Context, req *u.CreateUserRequest) (*u.SuccessResponse, error) {
	data := srv.transformUserRPC(req)
	if data.Email == "" || data.Name == "" {
		return &u.SuccessResponse{}, errors.New("please provide all fields")
	}

	if err := srv.useCase.Create(data); err != nil {
		return &u.SuccessResponse{}, err
	}
	return srv.transformUser(), nil
}

func (srv *UserServiceStruct) Read(ctx context.Context, req *u.SingleUserRequest) (*u.UserProfileResponse, error) {
	fmt.Println("Read user details")
	id := req.GetId()
	if id == "" {
		return &u.UserProfileResponse{}, errors.New("id cannot be blank")
	}

	user, err := srv.useCase.Read(id)
	if err != nil {
		return &u.UserProfileResponse{}, errors.New("user not found")
	}
	return srv.transformUserModel(user), nil
}
