// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Record struct {
	ID             string             `db:"id" json:"id"`
	OrganizationID string             `db:"organization_id" json:"organization_id"`
	UserID         string             `db:"user_id" json:"user_id"`
	Record         string             `db:"record" json:"record"`
	ScanTitle      string             `db:"scan_title" json:"scan_title"`
	CreatedAt      pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

type Upload struct {
	ID             string             `db:"id" json:"id"`
	OrganizationID string             `db:"organization_id" json:"organization_id"`
	UserID         string             `db:"user_id" json:"user_id"`
	ScanTitle      string             `db:"scan_title" json:"scan_title"`
	CreatedAt      pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}
