-- name: CreateRecord :one
INSERT INTO records (
  organization_id,
  user_id,
  record
) VALUES  (
  $1, $2, $3
) RETURNING *;

-- name: GetRecord :one
SELECT * FROM records where id = $1;

-- name: GetRecords :many
SELECT * FROM records where organization_id = $1
LIMIT 50;

-- name: UploadRecord :one
INSERT INTO uploads (
  organization_id,
  user_id,
  scan_title,
  file_name
) VALUES  (
  $1, $2, $3, $4
) RETURNING *; 
