-- name: CreateRecord :one
INSERT INTO records (
  diagnostic_id,
  user_id,
  record,
  scan_title
) VALUES  (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetRecord :one
SELECT * FROM records where id = $1;

-- name: GetRecords :many
SELECT * FROM records where diagnostic_id = $1
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
  diagnostic_id,
  user_id,
  scan_title
) VALUES  (
  $1, $2, $3
) RETURNING *; 

-- name: SearchRecordByNin :many
SELECT * FROM public.records JOIN public.users ON users.nin = $1;

-- name: SearchRecordByNinAndScanTitle :many
SELECT * FROM records JOIN public.users ON
users.nin = $1 WHERE scan_title ILIKE $2
LIMIT 10;
