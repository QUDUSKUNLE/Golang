// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateRecord(ctx context.Context, arg CreateRecordParams) (*Record, error)
	GetRecord(ctx context.Context, id string) (*Record, error)
	GetRecords(ctx context.Context, organizationID string) ([]*Record, error)
	GetRecordsByUser(ctx context.Context, userID string) ([]*Record, error)
	GetRecordsByUserAndScanTitle(ctx context.Context, arg GetRecordsByUserAndScanTitleParams) ([]*Record, error)
	UploadRecord(ctx context.Context, arg UploadRecordParams) (*Upload, error)
}

var _ Querier = (*Queries)(nil)
