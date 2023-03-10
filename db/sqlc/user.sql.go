// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username,
                   hashed_password,
                   full_name,
                   email)
VALUES ($1, $2, $3, $4) RETURNING username, hashed_password, full_name, email, password_updated_at, created_by, created_at, updated_by, updated_at, mark_for_delete
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordUpdatedAt,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedBy,
		&i.UpdatedAt,
		&i.MarkForDelete,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, full_name, email, password_updated_at, created_by, created_at, updated_by, updated_at, mark_for_delete
FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordUpdatedAt,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedBy,
		&i.UpdatedAt,
		&i.MarkForDelete,
	)
	return i, err
}
