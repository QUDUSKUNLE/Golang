-- name: GetUser :one
SELECT * FROM users where id = $1;

-- name: CreateUser :one
INSERT INTO users (
  email,
  nin,
  password,
  user_type
) VALUES  (
  $1, $2, $3, $4
) RETURNING id, email, nin, user_type, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateNin :one
UPDATE users
SET
  nin = COALESCE($1, nin),
  updated_at = NOW()
WHERE id = $2
RETURNING id, email, user_type, created_at, updated_at;
