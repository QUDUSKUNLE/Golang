package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/events"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	userProtoc "github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create a record
func (srv *UserServiceStruct) Create(ctx context.Context, req *userProtoc.CreateUserRequest) (*userProtoc.SuccessResponse, error) {
	// Transform request data
	data := srv.transformUserRPC(req)
	newUser, err := dto.BuildNewUser(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Create user in the database
	user, err := srv.userService.CreateUser(
		ctx, db.CreateUserParams{
			Email:    newUser.Email,
			Nin:      pgtype.Text{String: "", Valid: true},
			Password: newUser.Password,
			UserType: newUser.UserType,
		})
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	}
	// Check if user is an organization
	switch data.UserType {
	case db.UserEnumORGANIZATION:
		if err := srv.eventBroker.Publish(ctx, events.USER_EVENTS, &dto.UserCreatedEvent{UserID: user.ID, Email: newUser.Email.String}); err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		return &userProtoc.SuccessResponse{Data: OrganizationRegisteredSuccessfully}, nil
	case db.UserEnumDIAGNOSTIC:
		if err := srv.eventBroker.Publish(ctx, events.DIAGNOSTIC_EVENTS, &dto.UserCreatedEvent{UserID: user.ID, Email: newUser.Email.String}); err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		return &userProtoc.SuccessResponse{Data: OrganizationRegisteredSuccessfully}, nil
	default:
		return &userProtoc.SuccessResponse{Data: UserRegisteredSuccessfully}, nil
	}
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
	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	if limit <= 0 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	users, err := srv.userService.GetUsers(ctx, db.GetUsersParams{
		Limit: int32(limit), Offset: int32(offset)})
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
