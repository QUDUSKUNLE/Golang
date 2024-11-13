package handler

import (
	"github.com/QUDUSKUNLE/microservices/services/auth-service/internal/models"
	u "github.com/QUDUSKUNLE/microservices/services/auth-service/protogen/golang/user"
)

func (srv *UserServiceStruct) transformUserRPC(req *u.CreateUserRequest) models.User {
	return models.User{Name: req.GetName(), Email: req.GetEmail()}
}

func (srv *UserServiceStruct) transformUserModel(user models.User) *u.UserProfileResponse {
	return &u.UserProfileResponse{Name: user.Name, Email: user.Email}
}

func (srv *UserServiceStruct) transformUser() *u.SuccessResponse {
	return &u.SuccessResponse{Response: "User registered successfully"}
}
