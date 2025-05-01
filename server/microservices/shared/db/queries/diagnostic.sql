-- name: CreateDiagnostic :one
INSERT INTO diagnostics (
  user_id
) VALUES  (
  $1
) RETURNING *;

-- name: GetDiagnostic :one
SELECT * FROM diagnostics WHERE id = $1;
