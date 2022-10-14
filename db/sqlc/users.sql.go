// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username, email, password, firstname, lastname, phone
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, created_at, updated_at, username, email, password, firstname, lastname, phone
`

type CreateUserParams struct {
	Username  sql.NullString `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Phone     sql.NullString `json:"phone"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Firstname,
		arg.Lastname,
		arg.Phone,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Firstname,
		&i.Lastname,
		&i.Phone,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, username, email, password, firstname, lastname, phone FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Firstname,
		&i.Lastname,
		&i.Phone,
	)
	return i, err
}
