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
SELECT * FROM diagnostic_schedules WHERE id = $1;

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
  user_id = COALESCE($1, user_id),
  diagnostic_centre_id = COALESCE($2, diagnostic_centre_id),
  date = COALESCE($3, date),
  time = COALESCE($4, time),
  test_type = COALESCE($5, test_type),
  status = COALESCE($6, status),
  notes = COALESCE($7, notes),
  updated_at = NOW()
WHERE id = $8
RETURNING *;

-- name: DeleteSchedule :one
DELETE FROM diagnostic_schedules
WHERE id = $1
RETURNING *;

