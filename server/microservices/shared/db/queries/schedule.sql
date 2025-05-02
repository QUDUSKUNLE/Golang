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

-- name: GetSchedulesByCentre :many
SELECT * FROM diagnostic_schedules
WHERE diagnostic_centre_id = $1
ORDER BY date DESC, time DESC
LIMIT $2 OFFSET $3;

-- name: GetSchedulesByStatus :many
SELECT * FROM diagnostic_schedules
WHERE status = $1
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
WHERE id = $6
RETURNING *;

-- name: CancelSchedule :one
DELETE FROM diagnostic_schedules
WHERE id = $1 AND user_id = $2
RETURNING *;
