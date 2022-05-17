// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: members.sql

package db

import (
	"context"
)

const createMember = `-- name: CreateMember :one
INSERT INTO members (
  full_name, email, hashed_password, status_code
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, full_name, email, hashed_password, password_changed_at, created_at, status_code
`

type CreateMemberParams struct {
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	StatusCode     string `json:"status_code"`
}

func (q *Queries) CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, createMember,
		arg.FullName,
		arg.Email,
		arg.HashedPassword,
		arg.StatusCode,
	)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.StatusCode,
	)
	return i, err
}

const deleteMember = `-- name: DeleteMember :exec
DELETE FROM members
WHERE id = $1
`

func (q *Queries) DeleteMember(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteMember, id)
	return err
}

const getMember = `-- name: GetMember :one
SELECT id, full_name, email, hashed_password, password_changed_at, created_at, status_code FROM members
WHERE id = $1 
LIMIT 1
`

func (q *Queries) GetMember(ctx context.Context, id int32) (Member, error) {
	row := q.db.QueryRowContext(ctx, getMember, id)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.StatusCode,
	)
	return i, err
}

const getMemberByEmail = `-- name: GetMemberByEmail :one
SELECT id, full_name, email, hashed_password, password_changed_at, created_at, status_code FROM members
WHERE email = $1 
LIMIT 1
`

func (q *Queries) GetMemberByEmail(ctx context.Context, email string) (Member, error) {
	row := q.db.QueryRowContext(ctx, getMemberByEmail, email)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.StatusCode,
	)
	return i, err
}

const listMembers = `-- name: ListMembers :many
SELECT id, full_name, email, hashed_password, password_changed_at, created_at, status_code FROM members
ORDER BY created_at
`

func (q *Queries) ListMembers(ctx context.Context) ([]Member, error) {
	rows, err := q.db.QueryContext(ctx, listMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Member{}
	for rows.Next() {
		var i Member
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Email,
			&i.HashedPassword,
			&i.PasswordChangedAt,
			&i.CreatedAt,
			&i.StatusCode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMembersPaginated = `-- name: ListMembersPaginated :many
SELECT id, full_name, email, hashed_password, password_changed_at, created_at, status_code FROM members
ORDER BY created_at
LIMIT $1
OFFSET $2
`

type ListMembersPaginatedParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListMembersPaginated(ctx context.Context, arg ListMembersPaginatedParams) ([]Member, error) {
	rows, err := q.db.QueryContext(ctx, listMembersPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Member{}
	for rows.Next() {
		var i Member
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Email,
			&i.HashedPassword,
			&i.PasswordChangedAt,
			&i.CreatedAt,
			&i.StatusCode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMember = `-- name: UpdateMember :one
UPDATE members 
SET full_name = $2, email = $3, status_code = $4
WHERE id=$1
RETURNING id, full_name, email, hashed_password, password_changed_at, created_at, status_code
`

type UpdateMemberParams struct {
	ID         int32  `json:"id"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	StatusCode string `json:"status_code"`
}

func (q *Queries) UpdateMember(ctx context.Context, arg UpdateMemberParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, updateMember,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.StatusCode,
	)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.StatusCode,
	)
	return i, err
}

const updateMemberStaus = `-- name: UpdateMemberStaus :one
UPDATE members 
SET status_code = $2
WHERE id=$1
RETURNING id, full_name, email, hashed_password, password_changed_at, created_at, status_code
`

type UpdateMemberStausParams struct {
	ID         int32  `json:"id"`
	StatusCode string `json:"status_code"`
}

func (q *Queries) UpdateMemberStaus(ctx context.Context, arg UpdateMemberStausParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, updateMemberStaus, arg.ID, arg.StatusCode)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.StatusCode,
	)
	return i, err
}
