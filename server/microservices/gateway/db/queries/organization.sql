-- name: GetOrganization :one
SELECT * FROM organizations WHERE id = $1;

-- name: GetOrganizationByUserID :one
SELECT * FROM organizations WHERE user_id = $1;

-- name: CreateOrganization :one
INSERT INTO organizations (
  user_id
) VALUES  (
  $1
) RETURNING *;
