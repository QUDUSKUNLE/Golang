package handler

import (
	"context"
	"fmt"

	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/QUDUSKUNLE/microservices/shared/events"
	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	userProtoc "github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
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
		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("Error creating a new user account %s", err.Error()))
	}
	// Check if user is an organization
	switch data.UserType {
	case db.UserEnumORGANIZATION:
		if err := srv.publishUserCreatedEvent(ctx, events.USER_EVENTS, user.ID); err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		return &userProtoc.SuccessResponse{Data: OrganizationRegisteredSuccessfully}, nil
	case db.UserEnumDIAGNOSTIC:
		if err := srv.publishUserCreatedEvent(ctx, events.DIAGNOSTIC_EVENTS, user.ID); err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		if err := srv.publishNotificationEvent(ctx, user.ID, "Diagnostic SignUp", "Email", map[string]string{"email": user.Email.String}); err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		return &userProtoc.SuccessResponse{Data: DiagnosticRegisteredSuccessfully}, nil
	default:
		if err := srv.publishNotificationEvent(ctx, user.ID, "User SignUp", "Email", map[string]string{
			"email": user.Email.String,
		}); err != nil {
			return nil, status.Error(codes.Aborted, fmt.Sprintf("Failed to publish user notification: %v", err))
		}
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
	_, err := middleware.ValidateUser(ctx, string(db.UserEnumADMIN))
	if err != nil {
		return nil, err
	}
	limit, offset := utils.PaginationParams(req.GetLimit(), req.GetOffset())
	users, err := srv.userService.GetUsers(ctx, db.GetUsersParams{
		Limit: limit, Offset: offset})
	if err != nil {
		return nil, status.Error(codes.NotFound, NotFound)
	}
	usersResponse := &userProtoc.GetUsersResponse{Data: []*userProtoc.User{}}
	for _, user := range users {
		usersResponse.Data = append(usersResponse.Data, transformUserToProto(*user))
	}
	return usersResponse, nil
}

func (srv *UserServiceStruct) UpdateUser(ctx context.Context, req *userProtoc.UpdateUserRequest) (*userProtoc.UpdateUserResponse, error) {
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		return nil, err
	}
	addressJSON, err := utils.MapToStruct(req.GetAddress().AsMap())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid address format: %v", err)
	}
	contactJSON, err := utils.MapToStruct(req.GetContact().AsMap())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid contact format: %v", err)
	}
	address, _ := utils.JsonMarshal(addressJSON)
	contact, _ := utils.JsonMarshal(contactJSON)
	_, err = srv.userService.UpdateUser(ctx, db.UpdateUserParams{
		Nin:     pgtype.Text{String: req.GetNin(), Valid: true},
		Address: address,
		Contact: contact,
		ID:      user.UserID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user: %v", err)
	}
	return &userProtoc.UpdateUserResponse{Data: UserUpdatedSuccessfully}, nil
}

func (srv *UserServiceStruct) Home(ctx context.Context, req *userProtoc.HomeRequest) (*userProtoc.GetHomeResponse, error) {
	return &userProtoc.GetHomeResponse{Message: WelcomeHome}, nil
}

// Helper method to publish user-created events
func (srv *UserServiceStruct) publishUserCreatedEvent(ctx context.Context, topic string, userID string) error {
	if err := srv.eventBroker.Publish(ctx, topic, &dto.UserCreatedEvent{UserID: userID}); err != nil {
		logger.GetLogger().Error("Failed to publish user-created event", zap.String("topic", topic), zap.String("userID", userID), zap.Error(err))
		return err
	}
	logger.GetLogger().Info("User-created event published successfully", zap.String("topic", topic), zap.String("userID", userID))
	return nil
}

// Helper method to publish notification events
func (srv *UserServiceStruct) publishNotificationEvent(ctx context.Context, userID string, eventType string, channel string, data map[string]string) error {
	if err := srv.eventBroker.Publish(ctx, events.NOTIFICATION_EVENTS, &dto.NotificationEvent{
		UserID:    userID,
		EventType: eventType,
		Channel:   channel,
		Data:      data,
	}); err != nil {
		logger.GetLogger().Error("Failed to publish notification event", zap.String("eventType", eventType), zap.String("channel", channel), zap.Error(err))
		return err
	}
	logger.GetLogger().Info("Notification event published successfully", zap.String("eventType", eventType), zap.String("channel", channel))
	return nil
}
