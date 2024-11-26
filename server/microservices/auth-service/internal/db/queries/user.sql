-- name: GetUser :one
SELECT * FROM users where id = $1;
