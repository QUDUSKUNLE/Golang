package grpc

import (
	"context"
	"errors"

	pb "github.com/QUDUSKUNLE/microservices/proto"
)

func (srv *UserServiceStruct) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.SuccessResponse, error) {
	data := srv.transformUserRPC(req)
	if data.Email == "" || data.Name == "" {
		return &pb.SuccessResponse{}, errors.New("please provide all fields")
	}

	if err := srv.useCase.Create(data); err != nil {
		return &pb.SuccessResponse{}, err
	}
	return srv.transformUser(), nil
}

func (srv *UserServiceStruct) Get(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	id := req.GetId()
	if id == "" {
		return &pb.UserProfileResponse{}, errors.New("id cannot be blank")
	}

	user, err := srv.useCase.Get(id)
	if err != nil {
		return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}
