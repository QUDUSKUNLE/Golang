-- name: CreateRecord :one
INSERT INTO records (
  organization_id,
  user_id,
  record,
  scan_title
) VALUES  (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetRecord :one
SELECT * FROM records where id = $1;

-- name: GetRecords :many
SELECT * FROM records where organization_id = $1
LIMIT 50;

-- name: GetRecordsByUser :many
SELECT * FROM records where user_id = $1
ORDER BY created_at DESC
LIMIT 10;

-- name: GetRecordsByUserAndScanTitle :many
SELECT * FROM records where user_id = $1 and scan_title ILIKE $2
ORDER BY created_at DESC
LIMIT 10;

-- name: UploadRecord :one
INSERT INTO uploads (
  organization_id,
  user_id,
  scan_title
) VALUES  (
  $1, $2, $3
) RETURNING *; 
