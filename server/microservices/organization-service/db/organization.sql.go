// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: organization.sql

package db

import (
	"context"
)

const createOrganization = `-- name: CreateOrganization :one
INSERT INTO organizations (
  user_id
) VALUES  (
  $1
) RETURNING id, user_id, created_at, updated_at
`

func (q *Queries) CreateOrganization(ctx context.Context, userID string) (*Organization, error) {
	row := q.db.QueryRow(ctx, createOrganization, userID)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getOrganization = `-- name: GetOrganization :one
SELECT id, user_id, created_at, updated_at FROM organizations WHERE id = $1
`

func (q *Queries) GetOrganization(ctx context.Context, id string) (*Organization, error) {
	row := q.db.QueryRow(ctx, getOrganization, id)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}