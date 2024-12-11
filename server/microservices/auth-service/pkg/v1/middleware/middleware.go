package middleware

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	switch (info.FullMethod) {
	case UpdateNin:
		return urinaryHelper(ctx, req, handler)
	default:
		return handler(ctx, req)
	}
}

func urinaryHelper(ctx context.Context, req any, handler grpc.UnaryHandler) (any, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metatdata is not provided")
	}
	// extract token from the authorization header
	token := meta["authorization"]
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
	}
	id, err := validateToken(ctx, strings.Split(token[0], " ")[1]);
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	ctx = context.WithValue(ctx, "UserID", id)
	return handler(ctx, req)
}

func validateToken(_ context.Context, token string) (string, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id, ok := claims["id"].(string)
		if !ok {
			return "", errors.New("failed to extract id from claims")
		}
		return id, nil
	}
	return "", errors.New("invalid token")
}
