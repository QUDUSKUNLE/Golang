package middleware

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/auth"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/record"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/schedule"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
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
		// Delegate validation to specific functions
		if err := validateRequest(req); err != nil {
			return nil, err
		}
		// Handle authorization for specific methods
		if requiresAuthorization(info.FullMethod) {
			return authorizationHelper(ctx, req, handler)
		}
		return handler(ctx, req)
	}
}

func validateRequest(req interface{}) error {
	switch r := req.(type) {
	case *user.CreateUserRequest:
		return validateCreateUserRequest(r)
	case *user.SingleUserRequest:
		return validateSingleUserRequest(r)
	case *user.UpdateUserRequest:
		return validateUpdateUserRequest(r)
	case *auth.SignInRequest:
		return validateSignInRequest(r)
	case *record.ScanUploadRequest:
		return validateScanUploadRequest(r)
	case *schedule.ScheduleRequest:
		return validateScheduleRequest(r)
	case *diagnostic.CreateDiagnosticRequest:
		return validateCreateDiagnosticRequest(r)
	case *diagnostic.GetDiagnosticRequest:
		return validateGetDiagnosticRequest(r)
	default:
		return nil
	}
}

func validateCreateUserRequest(r *user.CreateUserRequest) error {
	if !ValidateEmail(r.Email) || r.Password == "" || string(r.UserType) == "" {
		return status.Errorf(codes.InvalidArgument, "Email, Password, and UserType cannot be empty")
	}
	if r.Password != r.ConfirmPassword {
		return status.Errorf(codes.InvalidArgument, "Password and ConfirmPassword must match")
	}
	return nil
}

func validateSingleUserRequest(r *user.SingleUserRequest) error {
	if !ValidateUUID(r.Id) {
		return status.Errorf(codes.InvalidArgument, "Invalid ID")
	}
	return nil
}

func validateUpdateUserRequest(r *user.UpdateUserRequest) error {
	if !ValidateNIN(r.Nin) {
		return status.Errorf(codes.InvalidArgument, "Invalid NIN")
	}
	return nil
}

func validateSignInRequest(r *auth.SignInRequest) error {
	if !ValidateEmail(r.Email) || r.Password == "" {
		return status.Errorf(codes.InvalidArgument, "Email and Password are required")
	}
	return nil
}

func validateScanUploadRequest(r *record.ScanUploadRequest) error {
	if r.ScanTitle == "" {
		return status.Errorf(codes.InvalidArgument, "ScanTitle is required")
	}
	if r.FileName == "" {
		return status.Errorf(codes.InvalidArgument, "FileName is required")
	}
	if !ValidateUUID(r.UserId) {
		return status.Errorf(codes.InvalidArgument, "Invalid UserId")
	}
	return nil
}

func validateScheduleRequest(r *schedule.ScheduleRequest) error {
	if r.DiagnosticCentreId == "" {
		return status.Errorf(codes.InvalidArgument, "DiagnosticCentreId is required")
	}
	if r.Date == "" {
		return status.Errorf(codes.InvalidArgument, "Date is required")
	}
	if r.Time == "" {
		return status.Errorf(codes.InvalidArgument, "Time is required")
	}
	return nil
}

func validateCreateDiagnosticRequest(r *diagnostic.CreateDiagnosticRequest) error {
	if r.DiagnosticCentreName == "" {
		return status.Errorf(codes.InvalidArgument, "DiagnosticCentreId is required")
	}
	return nil
}

func validateGetDiagnosticRequest(r *diagnostic.GetDiagnosticRequest) error {
	if r.DiagnosticId == "" {
		return status.Errorf(codes.InvalidArgument, "DiagnosticId is required")
	}
	if !ValidateUUID(r.DiagnosticId) {
		return status.Errorf(codes.InvalidArgument, "Invalid DiagnosticId")
	}
	return nil
}

func requiresAuthorization(method string) bool {
	authorizedMethods := map[string]bool{
		constants.UpdateUser:                    true,
		constants.ReadUsers:                     true,
		constants.GetRecords:                    true,
		constants.GetRecord:                     true,
		constants.ScanUpload:                    true,
		constants.SearchRecord:                  true,
		constants.SearchByNin:                   true,
		constants.CreateDiagnostic:              true,
		constants.ListDiagnostics:               true,
		constants.GetDiagnostic:                 true,
		constants.UpdateDiagnostic:              true,
		constants.DeleteDiagnostic:              true,
		constants.CreateSchedule:                true,
		constants.GetSchedule:                   true,
		constants.ListSchedules:                 true,
		constants.CancelSchedule:                true,
		constants.UpdateSchedule:                true,
		constants.ListDiagnosticCentreSchedules: true,
		constants.ListDiagnosticSchedules:       true,
	}
	return authorizedMethods[method]
}

func authorizationHelper(ctx context.Context, req interface{}, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Metadata is not provided")
	}

	// Extract token from the authorization header
	token := meta["authorization"]
	if len(token) <= 1 {
		return nil, status.Error(codes.Unauthenticated, "Authorization token is not provided")
	}

	// Validate the token
	userType, err := validateToken(strings.Split(token[0], " ")[1])
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// Add user and token to the context
	ctx = context.WithValue(ctx, "user", userType)
	ctx = context.WithValue(ctx, "token", token[0])

	// Proceed to the next handler
	return handler(ctx, req)
}

func validateToken(token string) (*constants.UserType, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, errors.New("invalid token claims")
	}

	id, ok := claims["id"].(string)
	if !ok {
		return nil, errors.New("failed to extract id from claims")
	}

	userType, ok := claims["user_type"].(string)
	if !ok {
		return nil, errors.New("failed to extract user_type from claims")
	}

	return &constants.UserType{UserID: id, Type: userType}, nil
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

func GetUserFromContext(ctx context.Context) (*constants.UserType, error) {
	user, ok := ctx.Value("user").(*constants.UserType)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, constants.ErrUnauthorized)
	}
	return user, nil
}
