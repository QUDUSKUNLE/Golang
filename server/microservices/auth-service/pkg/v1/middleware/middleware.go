package middleware

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ValidationInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if r, ok := req.(*user.CreateUserRequest); ok {
			if !ValidateEmail(r.Email) || r.Password == "" || string(r.UserType) == "" {
				return nil, status.Errorf(codes.InvalidArgument, "Email, Password and UserType cannot be empty")
			} else if r.Password != r.ConfirmPassword {
				return nil, status.Errorf(codes.InvalidArgument, "Password and ConfirmPassword must match")
			}
		}
		if r, ok := req.(*user.SingleUserRequest); ok {
			if !ValidateUUID(r.Id) {
				return nil, status.Errorf(codes.InvalidArgument, "Id cannot be empty")
			}
		}
		if r, ok := req.(*user.SignInRequest); ok {
			if !ValidateEmail(r.Email) || r.Password == "" {
				return nil, status.Errorf(codes.InvalidArgument, "Email and Password cannot be empty")
			}
		}
		if r, ok := req.(*user.UpdateNinRequest); ok {
			if !ValidateNIN(r.Nin) {
				return nil, status.Errorf(codes.InvalidArgument, "Invalid NIN")
			}
		}
		switch info.FullMethod {
		case UpdateNin:
			return urinaryHelper(ctx, req, handler)
		default:
			return handler(ctx, req)
		}
	}
}

func urinaryHelper(ctx context.Context, req interface{}, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Metatdata is not provided")
	}
	// extract token from the authorization header
	token := meta["authorization"]
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Authorization token is not provided")
	}
	user_type, err := validateToken(ctx, strings.Split(token[0], " ")[1])
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	ctx = context.WithValue(ctx, "user", user_type)
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
		typ, ok := claims["user_type"].(string)
		if !ok {
			return &UserType{}, errors.New("failed to extract user_type from claims")
		}
		return &UserType{UserID: id, Type: typ}, nil
	}
	return &UserType{}, errors.New("invalid token")
}

func ValidateUUID(input string) bool {
	_, err := uuid.Parse(input)
	return err == nil
}

func ValidateEmail(email string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

func ValidateNIN(nin string) bool {
	return regexp.MustCompile(`^\d{11}$`).MatchString(nin)
}
