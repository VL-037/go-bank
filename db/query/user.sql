-- name: CreateUser :one
INSERT INTO users (owner,
                   balance,
                   currency)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetAccount :one
SELECT *
FROM users
WHERE id = $1 LIMIT 1;