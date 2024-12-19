package middleware

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ValidationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if r, ok := req.(*organization.CreateOrganizationRequest); ok {
		if r.UserId == "" {
			return nil, status.Errorf(codes.InvalidArgument, "User ID cannot be empty")
		}
	}
	if r, ok := req.(*organization.GetOrganizationRequest); ok {
		if r.Id == "" {
			return nil, status.Errorf(codes.InvalidArgument, "id is required")
		}
	}
	return handler(ctx, req)
}

func urinaryHelper(ctx context.Context, req interface{}, handler grpc.UnaryHandler) (any, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Metatdata is not provided")
	}
	// extract token from the authorization header
	token := meta["authorization"]
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Authorization token is not provided")
	}
	user, err := validateToken(ctx, strings.Split(token[0], " ")[1])
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	ctx = context.WithValue(ctx, "user", user)
	return handler(ctx, req)
}

func validateToken(_ context.Context, token string) (*UserType, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return &UserType{}, err
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id, ok := claims["id"].(string)
		if !ok {
			return &UserType{}, errors.New("failed to extract id from claims")
		}
		ty, ok := claims["user_type"].(string)
		if !ok {
			return &UserType{}, errors.New("failed to extract user_type from claims")
		}
		return &UserType{UserID: id, Type: ty}, nil
	}
	return &UserType{}, errors.New("failed to extract invalid token")
}
