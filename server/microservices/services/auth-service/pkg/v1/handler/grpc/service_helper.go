package grpc

import (
	"github.com/QUDUSKUNLE/microservices/services/auth-service/internal/models"
	g "github.com/QUDUSKUNLE/microservices/services/auth-service/protogen/golang/greet"
	u "github.com/QUDUSKUNLE/microservices/services/auth-service/protogen/golang/user"
)

func (srv *UserServiceStruct) transformMessage(req *g.GreetRequest) *g.GreetResponse {
	return &g.GreetResponse{Message: req.GetName()}
}

func (srv *UserServiceStruct) transformUserRPC(req *u.CreateUserRequest) models.User {
	return models.User{Name: req.GetName(), Email: req.GetEmail()}
}

func (srv *UserServiceStruct) transformUserModel(user models.User) *u.UserProfileResponse {
	return &u.UserProfileResponse{Name: user.Name, Email: user.Email}
}

func (srv *UserServiceStruct) transformUser() *u.SuccessResponse {
	return &u.SuccessResponse{Response: ""}
}
