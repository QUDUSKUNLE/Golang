-- name: CreateSchedule :one
INSERT INTO diagnostic_schedules (
  user_id,
  diagnostic_centre_id,
  date,
  time,
  test_type,
  status,
  notes
) VALUES  (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSchedule :one
SELECT * FROM diagnostic_schedules WHERE id = $1 AND user_id=$2;

-- name: GetSchedules :many
SELECT * FROM diagnostic_schedules
WHERE user_id = $1
ORDER BY date DESC, time DESC
LIMIT $2 OFFSET $3;

-- name: UpdateSchedule :one
UPDATE diagnostic_schedules
SET
  date = COALESCE($1, date),
  time = COALESCE($2, time),
  test_type = COALESCE($3, test_type),
  status = COALESCE($4, status),
  notes = COALESCE($5, notes),
  updated_at = NOW()
WHERE id = $6 AND user_id = $7
RETURNING *;

-- name: CancelSchedule :one
DELETE FROM diagnostic_schedules
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: GetSchedulesByCentre :many
SELECT * FROM diagnostic_schedules
WHERE diagnostic_centre_id = $1
ORDER BY date DESC, time DESC
LIMIT $2 OFFSET $3;

-- name: GetSchedulesByDiagnosticCentre :many
SELECT id, user_id, diagnostic_centre_id, date, time, test_type, status, notes, created_at, updated_at
FROM diagnostic_schedules
WHERE diagnostic_centre_id = $1
  AND ($2::schedule_status IS NULL OR status = $2)
  AND ($3::timestamp IS NULL OR date >= $3)
ORDER BY date DESC, time DESC
LIMIT $4 OFFSET $5;

-- name: GetScheduleByDiagnosticCentre :one
SELECT * FROM diagnostic_schedules
WHERE id = $1 AND diagnostic_centre_id = $2;

-- name: GetSchedulesDiagnosticCentreByStatusAndDate :many
SELECT id, user_id, diagnostic_centre_id, date, time, test_type, status, notes, created_at, updated_at
FROM diagnostic_schedules
WHERE diagnostic_centre_id = $1
  AND status = $2
  AND date >= $3 -- Filter schedules starting from a specific date
ORDER BY date DESC, time DESC
LIMIT $4 OFFSET $5;
