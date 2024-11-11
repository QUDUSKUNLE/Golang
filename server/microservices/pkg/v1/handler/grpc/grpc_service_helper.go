package grpc

import (
	"github.com/QUDUSKUNLE/microservices/internal/models"
	pb "github.com/QUDUSKUNLE/microservices/proto"
)

func (srv *UserServiceStruct) transformMessage(req *pb.GreetRequest) *pb.GreetResponse {
	return &pb.GreetResponse{Message: req.GetName()}
}

func (srv *UserServiceStruct) transformUserRPC(req *pb.CreateUserRequest) models.User{
  return models.User{Name: req.GetName(), Email: req.GetEmail()}
}

func (srv *UserServiceStruct) transformUserModel(user models.User) *pb.UserProfileResponse {
  return &pb.UserProfileResponse{Name: user.Name, Email: user.Email}
}

func (srv *UserServiceStruct) transformUser() *pb.SuccessResponse {
  return &pb.SuccessResponse{Response: ""}
}
