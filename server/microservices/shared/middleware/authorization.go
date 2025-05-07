package middleware

import (
	"context"
	"fmt"

	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateUser(ctx context.Context, userType string) (*constants.UserType, error) {
	user, ok := ctx.Value("user").(*constants.UserType)
	fmt.Println(user, "**********")
	if !ok || user.Type != userType {
		return nil, status.Errorf(codes.PermissionDenied, constants.ErrUnauthorized)
	}
	return user, nil
}

func AuthorizationMiddleware(requiredRole string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		_, err := ValidateUser(ctx, requiredRole)
		if err != nil {
			return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
		}
		return handler(ctx, req)
	}
}
