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
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
