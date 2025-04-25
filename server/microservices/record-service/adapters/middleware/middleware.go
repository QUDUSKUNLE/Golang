package middleware

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/QUDUSKUNLE/microservices/shared/protogen/record"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ValidateUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if r, ok := req.(*record.ScanUploadRequest); ok {
			if r.ScanTitle == "" || r.FileName == "" || !ValidateUUID(r.UserId) {
				return nil, status.Errorf(codes.InvalidArgument, "ScanTitle or FileName or UserID cannot be empty")
			}
		}
		switch info.FullMethod {
		case GetRecords, GetRecord, ScanUpload, SearchRecord, SearchByNin:
			return urinaryHelper(ctx, req, handler)
		default:
			return handler(ctx, req)
		}
	}
}

func urinaryHelper(
	ctx context.Context,
	req interface{},
	handler grpc.UnaryHandler,
) (interface{}, error) {
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
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
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
		typ, ok := claims["user_type"].(string)
		if !ok {
			return &UserType{}, errors.New("failed to extract type from claims")
		}
		return &UserType{UserID: id, Type: typ}, nil
	}
	return &UserType{}, errors.New("failed to extract invalid token")
}

func ValidateUUID(input string) bool {
	_, err := uuid.Parse(input)
	return err == nil
}
