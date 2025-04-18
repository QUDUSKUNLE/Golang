// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: records.sql

package db

import (
	"context"
)

const createRecord = `-- name: CreateRecord :one
INSERT INTO records (
  organization_id,
  user_id,
  record,
  scan_title
) VALUES  (
  $1, $2, $3, $4
) RETURNING id, organization_id, user_id, record, scan_title, created_at, updated_at
`

type CreateRecordParams struct {
	OrganizationID string `db:"organization_id" json:"organization_id"`
	UserID         string `db:"user_id" json:"user_id"`
	Record         string `db:"record" json:"record"`
	ScanTitle      string `db:"scan_title" json:"scan_title"`
}

func (q *Queries) CreateRecord(ctx context.Context, arg CreateRecordParams) (*Record, error) {
	row := q.db.QueryRow(ctx, createRecord,
		arg.OrganizationID,
		arg.UserID,
		arg.Record,
		arg.ScanTitle,
	)
	var i Record
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.UserID,
		&i.Record,
		&i.ScanTitle,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getRecord = `-- name: GetRecord :one
SELECT id, organization_id, user_id, record, scan_title, created_at, updated_at FROM records where id = $1
`

func (q *Queries) GetRecord(ctx context.Context, id string) (*Record, error) {
	row := q.db.QueryRow(ctx, getRecord, id)
	var i Record
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.UserID,
		&i.Record,
		&i.ScanTitle,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getRecords = `-- name: GetRecords :many
SELECT id, organization_id, user_id, record, scan_title, created_at, updated_at FROM records where organization_id = $1
LIMIT 50
`

func (q *Queries) GetRecords(ctx context.Context, organizationID string) ([]*Record, error) {
	rows, err := q.db.Query(ctx, getRecords, organizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Record
	for rows.Next() {
		var i Record
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.UserID,
			&i.Record,
			&i.ScanTitle,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRecordsByUser = `-- name: GetRecordsByUser :many
SELECT id, organization_id, user_id, record, scan_title, created_at, updated_at FROM records where user_id = $1
ORDER BY created_at DESC
LIMIT 1
`

func (q *Queries) GetRecordsByUser(ctx context.Context, userID string) ([]*Record, error) {
	rows, err := q.db.Query(ctx, getRecordsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Record
	for rows.Next() {
		var i Record
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.UserID,
			&i.Record,
			&i.ScanTitle,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRecordsByUserAndScanTitle = `-- name: GetRecordsByUserAndScanTitle :many
SELECT id, organization_id, user_id, record, scan_title, created_at, updated_at FROM records where user_id = $1 and scan_title = $2
ORDER BY created_at DESC
LIMIT 10
`

type GetRecordsByUserAndScanTitleParams struct {
	UserID    string `db:"user_id" json:"user_id"`
	ScanTitle string `db:"scan_title" json:"scan_title"`
}

func (q *Queries) GetRecordsByUserAndScanTitle(ctx context.Context, arg GetRecordsByUserAndScanTitleParams) ([]*Record, error) {
	rows, err := q.db.Query(ctx, getRecordsByUserAndScanTitle, arg.UserID, arg.ScanTitle)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Record
	for rows.Next() {
		var i Record
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.UserID,
			&i.Record,
			&i.ScanTitle,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const uploadRecord = `-- name: UploadRecord :one
INSERT INTO uploads (
  organization_id,
  user_id,
  scan_title
) VALUES  (
  $1, $2, $3
) RETURNING id, organization_id, user_id, scan_title, created_at, updated_at
`

type UploadRecordParams struct {
	OrganizationID string `db:"organization_id" json:"organization_id"`
	UserID         string `db:"user_id" json:"user_id"`
	ScanTitle      string `db:"scan_title" json:"scan_title"`
}

func (q *Queries) UploadRecord(ctx context.Context, arg UploadRecordParams) (*Upload, error) {
	row := q.db.QueryRow(ctx, uploadRecord, arg.OrganizationID, arg.UserID, arg.ScanTitle)
	var i Upload
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.UserID,
		&i.ScanTitle,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
