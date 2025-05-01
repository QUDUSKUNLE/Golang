package middleware

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateUser(ctx context.Context, userType string) (*constants.UserType, error) {
	user, ok := ctx.Value("user").(*constants.UserType)
	if !ok || user.Type != userType {
		return nil, status.Errorf(codes.Unauthenticated, constants.ErrUnauthorized)
	}
	return user, nil
}
