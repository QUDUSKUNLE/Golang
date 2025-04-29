package handler

import (
	"context"
	"fmt"

	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	userProtoc "github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create a record
func (srv *UserServiceStruct) Create(ctx context.Context, req *userProtoc.CreateUserRequest) (*userProtoc.SuccessResponse, error) {

	fmt.Println("Hereeeee")
	// Transform request data
	data := srv.transformUserRPC(req)
	built_user, err := dto.BuildNewUser(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Create user in the database
	_, err = srv.userService.CreateUser(
		ctx, db.CreateUserParams{
			Email:    built_user.Email,
			Nin:      pgtype.Text{String: "", Valid: true},
			Password: built_user.Password,
			UserType: built_user.UserType,
		})
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	}
	// Check if user is an organization
	if data.UserType != db.UserEnumUSER {
		// if err := srv.eventBroker.Publish("CreatedUser", &dto.UserCreatedEvent{UserID: user.ID, Email: built_user.Email.String}); err != nil {
		// 	return nil, status.Error(codes.Aborted, err.Error())
		// }
		return &userProtoc.SuccessResponse{Data: OrganizationRegisteredSuccessfully}, nil
	}
	return &userProtoc.SuccessResponse{Data: UserRegisteredSuccessfully}, nil
}

func (srv *UserServiceStruct) Read(ctx context.Context, req *userProtoc.SingleUserRequest) (*userProtoc.GetUserResponse, error) {
	// Get a user with the ID
	user, err := srv.userService.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, NotFound)
	}
	data := transformUserToProto(*user)
	return &userProtoc.GetUserResponse{
		Data: data}, nil
}

func (srv *UserServiceStruct) ReadUsers(ctx context.Context, req *userProtoc.GetUsersRequest) (*userProtoc.GetUsersResponse, error) {
	// Check if user has admin right
	admin, err := getUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if admin.Type != string(db.UserEnumADMIN) {
		return nil, status.Errorf(codes.Unauthenticated, ErrUnauthorized)
	}
	users, err := srv.userService.GetUsers(ctx)
	if err != nil {
		return nil, status.Error(codes.NotFound, NotFound)
	}
	usersResponse := &userProtoc.GetUsersResponse{Data: []*userProtoc.User{}}
	for _, user := range users {
		usersResponse.Data = append(usersResponse.Data, transformUserToProto(*user))
	}
	return usersResponse, nil
}

func (srv *UserServiceStruct) UpdateNin(ctx context.Context, req *userProtoc.UpdateNinRequest) (*userProtoc.UpdateNinResponse, error) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	_, err = srv.userService.UpdateNin(ctx, db.UpdateNinParams{
		Nin: pgtype.Text{
			String: req.GetNin(), Valid: true,
		},
		ID: user.UserID})
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, err.Error())
	}
	return &userProtoc.UpdateNinResponse{Data: NinUpdatedSuccessfully}, nil
}

func (srv *UserServiceStruct) Home(ctx context.Context, req *userProtoc.HomeRequest) (*userProtoc.GetHomeResponse, error) {
	return &userProtoc.GetHomeResponse{Message: WelcomeHome}, nil
}
