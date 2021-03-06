// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: statuses.sql

package db

import (
	"context"
)

const createStatus = `-- name: CreateStatus :one
INSERT INTO statuses (
  code, name
) VALUES (
  $1, $2
)
RETURNING code, name, created_at
`

type CreateStatusParams struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (q *Queries) CreateStatus(ctx context.Context, arg CreateStatusParams) (Status, error) {
	row := q.db.QueryRowContext(ctx, createStatus, arg.Code, arg.Name)
	var i Status
	err := row.Scan(&i.Code, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteStatus = `-- name: DeleteStatus :exec
DELETE FROM statuses
WHERE code = $1
`

func (q *Queries) DeleteStatus(ctx context.Context, code string) error {
	_, err := q.db.ExecContext(ctx, deleteStatus, code)
	return err
}

const getStatus = `-- name: GetStatus :one
SELECT code, name, created_at FROM statuses
WHERE code = $1 LIMIT 1
`

func (q *Queries) GetStatus(ctx context.Context, code string) (Status, error) {
	row := q.db.QueryRowContext(ctx, getStatus, code)
	var i Status
	err := row.Scan(&i.Code, &i.Name, &i.CreatedAt)
	return i, err
}

const listStatuses = `-- name: ListStatuses :many
SELECT code, name, created_at FROM statuses
ORDER BY name
`

func (q *Queries) ListStatuses(ctx context.Context) ([]Status, error) {
	rows, err := q.db.QueryContext(ctx, listStatuses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Status{}
	for rows.Next() {
		var i Status
		if err := rows.Scan(&i.Code, &i.Name, &i.CreatedAt); err != nil {
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

const listStatusesPaginated = `-- name: ListStatusesPaginated :many
SELECT code, name, created_at FROM statuses
ORDER BY name
LIMIT $1
OFFSET $2
`

type ListStatusesPaginatedParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStatusesPaginated(ctx context.Context, arg ListStatusesPaginatedParams) ([]Status, error) {
	rows, err := q.db.QueryContext(ctx, listStatusesPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Status{}
	for rows.Next() {
		var i Status
		if err := rows.Scan(&i.Code, &i.Name, &i.CreatedAt); err != nil {
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

const updateStatus = `-- name: UpdateStatus :one
UPDATE statuses 
SET name = $2 
WHERE code=$1
RETURNING code, name, created_at
`

type UpdateStatusParams struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (q *Queries) UpdateStatus(ctx context.Context, arg UpdateStatusParams) (Status, error) {
	row := q.db.QueryRowContext(ctx, updateStatus, arg.Code, arg.Name)
	var i Status
	err := row.Scan(&i.Code, &i.Name, &i.CreatedAt)
	return i, err
}
