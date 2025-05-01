package services

import (
	"context"
	"time"

	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func parseTimestamp(input string) pgtype.Timestamptz {
	parsedTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return pgtype.Timestamptz{Valid: false} // Handle invalid date format
	}
	return pgtype.Timestamptz{Time: parsedTime, Valid: true}
}

func validateUser(ctx context.Context, userType string) (*constants.UserType, error) {
	user, ok := ctx.Value("user").(*constants.UserType)
	if !ok || user.Type != userType {
		return nil, status.Errorf(codes.Unauthenticated, constants.ErrUnauthorized)
	}
	return user, nil
}
