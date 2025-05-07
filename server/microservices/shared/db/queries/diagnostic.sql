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

-- name: GetDiagnosticCentreSchedules :many
SELECT 
  ds.id AS schedule_id,
  ds.user_id,
  ds.date,
  ds.time,
  ds.test_type,
  ds.status,
  ds.notes,
  d.id AS diagnostic_id,
  d.diagnostic_centre_name
FROM diagnostic_schedules ds
JOIN diagnostics d ON ds.diagnostic_centre_id = d.id
WHERE d.id = $1
ORDER BY ds.date DESC, ds.time DESC
LIMIT $2 OFFSET $3;

-- name: GetDiagnosticCentreUpcomingSchedules :many
SELECT 
  ds.id AS schedule_id,
  ds.user_id,
  ds.date,
  ds.time,
  ds.test_type,
  ds.status,
  ds.notes,
  d.diagnostic_centre_name
FROM diagnostic_schedules ds
JOIN diagnostics d ON ds.diagnostic_centre_id = d.id
WHERE d.id = $1 AND ds.date >= NOW()
ORDER BY ds.date ASC, ds.time ASC
LIMIT $2 OFFSET $3;

-- name: GetDiagnosticCentreSchedulesByTestType :many
SELECT 
  ds.id AS schedule_id,
  ds.user_id,
  ds.date,
  ds.time,
  ds.test_type,
  ds.status,
  ds.notes,
  d.diagnostic_centre_name
FROM diagnostic_schedules ds
JOIN diagnostics d ON ds.diagnostic_centre_id = d.id
WHERE d.id = $1 AND ds.test_type ILIKE '%' || $2 || '%'
ORDER BY ds.date DESC, ds.time DESC
LIMIT $3 OFFSET $4;

-- name: GetDiagnosticCentreSchedulesByStatus :many
SELECT 
  ds.id AS schedule_id,
  ds.user_id,
  ds.date,
  ds.time,
  ds.test_type,
  ds.status,
  ds.notes,
  d.diagnostic_centre_name
FROM diagnostic_schedules ds
JOIN diagnostics d ON ds.diagnostic_centre_id = d.id
WHERE d.id = $1 AND ds.status = $2
ORDER BY ds.date DESC, ds.time DESC
LIMIT $3 OFFSET $4;

-- name: GetDiagnosticCentreSchedulesBySpecificDate :many
SELECT 
  ds.id AS schedule_id,
  ds.user_id,
  ds.date,
  ds.time,
  ds.test_type,
  ds.status,
  ds.notes,
  d.diagnostic_centre_name
FROM diagnostic_schedules ds
JOIN diagnostics d ON ds.diagnostic_centre_id = d.id
WHERE d.id = $1 AND ds.date::DATE = $2
ORDER BY ds.time ASC
LIMIT $3 OFFSET $4;

-- name: GetDiagnosticCentreSchedulesWithDiagnosticsDetails :many
SELECT 
  ds.id AS schedule_id,
  ds.user_id,
  ds.date,
  ds.time,
  ds.test_type,
  ds.status,
  ds.notes,
  d.diagnostic_centre_name,
  d.latitude,
  d.longitude,
  d.address,
  d.contact
FROM diagnostic_schedules ds
JOIN diagnostics d ON ds.diagnostic_centre_id = d.id
WHERE d.id = $1
ORDER BY ds.date DESC, ds.time DESC
LIMIT $2 OFFSET $3;
