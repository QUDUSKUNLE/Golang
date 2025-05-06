-- name: CreateDiagnostic :one
-- Inserts a new diagnostic record into the diagnostics table.
INSERT INTO diagnostics (
  user_id,
  diagnostic_centre_name
) VALUES  (
  $1, $2
) RETURNING *;

-- name: GetDiagnostic :one
-- Retrieves a single diagnostic record by its ID.
SELECT * FROM diagnostics WHERE id = $1;

-- name: GetAllDiagnostics :many
-- Retrieves all diagnostic records with pagination.
SELECT * FROM diagnostics
ORDER BY created_at DESC
LIMIT $1 OFFSET $2; 

-- name: UpdateDiagnostic :one
-- Updates a diagnostic record by its ID.
UPDATE diagnostics
SET
  diagnostic_centre_name = COALESCE($2, diagnostic_centre_name),
  latitude = COALESCE($3, latitude),
  longitude = COALESCE($4, longitude),
  address = COALESCE($5, address),
  contact = COALESCE($6, contact),
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteDiagnostic :one
-- Deletes a diagnostic record by its ID.
DELETE FROM diagnostics WHERE id = $1 RETURNING *;

-- name: ListDiagnostics :many
-- Retrieves all diagnostic records for a specific user.
SELECT * FROM diagnostics WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: SearchDiagnostics :many
-- Searches diagnostics by name with pagination.
SELECT *
FROM diagnostics
WHERE name ILIKE '%' || $1 || '%'
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;
