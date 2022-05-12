// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: deposits.sql

package db

import (
	"context"
	"database/sql"
)

const createDeposit = `-- name: CreateDeposit :one
INSERT INTO deposits (
  reference_num, amount, currency, member_id, status_code
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, reference_num, amount, currency, member_id, created_at, status_code
`

type CreateDepositParams struct {
	ReferenceNum string         `json:"reference_num"`
	Amount       int64          `json:"amount"`
	Currency     sql.NullString `json:"currency"`
	MemberID     int32          `json:"member_id"`
	StatusCode   string         `json:"status_code"`
}

func (q *Queries) CreateDeposit(ctx context.Context, arg CreateDepositParams) (Deposit, error) {
	row := q.db.QueryRowContext(ctx, createDeposit,
		arg.ReferenceNum,
		arg.Amount,
		arg.Currency,
		arg.MemberID,
		arg.StatusCode,
	)
	var i Deposit
	err := row.Scan(
		&i.ID,
		&i.ReferenceNum,
		&i.Amount,
		&i.Currency,
		&i.MemberID,
		&i.CreatedAt,
		&i.StatusCode,
	)
	return i, err
}

const deleteDeposit = `-- name: DeleteDeposit :exec
DELETE FROM deposits
WHERE id = $1
`

func (q *Queries) DeleteDeposit(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteDeposit, id)
	return err
}

const getDeposit = `-- name: GetDeposit :one
SELECT id, reference_num, amount, currency, member_id, created_at, status_code FROM deposits
WHERE id = $1 
LIMIT 1
`

func (q *Queries) GetDeposit(ctx context.Context, id int32) (Deposit, error) {
	row := q.db.QueryRowContext(ctx, getDeposit, id)
	var i Deposit
	err := row.Scan(
		&i.ID,
		&i.ReferenceNum,
		&i.Amount,
		&i.Currency,
		&i.MemberID,
		&i.CreatedAt,
		&i.StatusCode,
	)
	return i, err
}

const getDepositByReferenceNum = `-- name: GetDepositByReferenceNum :one
SELECT id, reference_num, amount, currency, member_id, created_at, status_code FROM deposits
WHERE reference_num = $1 
LIMIT 1
`

func (q *Queries) GetDepositByReferenceNum(ctx context.Context, referenceNum string) (Deposit, error) {
	row := q.db.QueryRowContext(ctx, getDepositByReferenceNum, referenceNum)
	var i Deposit
	err := row.Scan(
		&i.ID,
		&i.ReferenceNum,
		&i.Amount,
		&i.Currency,
		&i.MemberID,
		&i.CreatedAt,
		&i.StatusCode,
	)
	return i, err
}

const listDeposits = `-- name: ListDeposits :many
SELECT id, reference_num, amount, currency, member_id, created_at, status_code FROM deposits
ORDER BY created_at
`

func (q *Queries) ListDeposits(ctx context.Context) ([]Deposit, error) {
	rows, err := q.db.QueryContext(ctx, listDeposits)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Deposit{}
	for rows.Next() {
		var i Deposit
		if err := rows.Scan(
			&i.ID,
			&i.ReferenceNum,
			&i.Amount,
			&i.Currency,
			&i.MemberID,
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

const listDepositsByMember = `-- name: ListDepositsByMember :many
SELECT id, reference_num, amount, currency, member_id, created_at, status_code FROM deposits
WHERE member_id = $1
ORDER BY created_at
`

func (q *Queries) ListDepositsByMember(ctx context.Context, memberID int32) ([]Deposit, error) {
	rows, err := q.db.QueryContext(ctx, listDepositsByMember, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Deposit{}
	for rows.Next() {
		var i Deposit
		if err := rows.Scan(
			&i.ID,
			&i.ReferenceNum,
			&i.Amount,
			&i.Currency,
			&i.MemberID,
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

const listDepositsPaginated = `-- name: ListDepositsPaginated :many
SELECT id, reference_num, amount, currency, member_id, created_at, status_code FROM deposits
ORDER BY created_at
LIMIT $1
OFFSET $2
`

type ListDepositsPaginatedParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListDepositsPaginated(ctx context.Context, arg ListDepositsPaginatedParams) ([]Deposit, error) {
	rows, err := q.db.QueryContext(ctx, listDepositsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Deposit{}
	for rows.Next() {
		var i Deposit
		if err := rows.Scan(
			&i.ID,
			&i.ReferenceNum,
			&i.Amount,
			&i.Currency,
			&i.MemberID,
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

const updateDeposit = `-- name: UpdateDeposit :one
UPDATE deposits 
SET reference_num = $2, amount = $3, currency = $4, member_id = $5, status_code = $6
WHERE id=$1
RETURNING id, reference_num, amount, currency, member_id, created_at, status_code
`

type UpdateDepositParams struct {
	ID           int32          `json:"id"`
	ReferenceNum string         `json:"reference_num"`
	Amount       int64          `json:"amount"`
	Currency     sql.NullString `json:"currency"`
	MemberID     int32          `json:"member_id"`
	StatusCode   string         `json:"status_code"`
}

func (q *Queries) UpdateDeposit(ctx context.Context, arg UpdateDepositParams) (Deposit, error) {
	row := q.db.QueryRowContext(ctx, updateDeposit,
		arg.ID,
		arg.ReferenceNum,
		arg.Amount,
		arg.Currency,
		arg.MemberID,
		arg.StatusCode,
	)
	var i Deposit
	err := row.Scan(
		&i.ID,
		&i.ReferenceNum,
		&i.Amount,
		&i.Currency,
		&i.MemberID,
		&i.CreatedAt,
		&i.StatusCode,
	)
	return i, err
}

const updateDepositStaus = `-- name: UpdateDepositStaus :one
UPDATE deposits 
SET status_code = $2
WHERE id=$1
RETURNING id, reference_num, amount, currency, member_id, created_at, status_code
`

type UpdateDepositStausParams struct {
	ID         int32  `json:"id"`
	StatusCode string `json:"status_code"`
}

func (q *Queries) UpdateDepositStaus(ctx context.Context, arg UpdateDepositStausParams) (Deposit, error) {
	row := q.db.QueryRowContext(ctx, updateDepositStaus, arg.ID, arg.StatusCode)
	var i Deposit
	err := row.Scan(
		&i.ID,
		&i.ReferenceNum,
		&i.Amount,
		&i.Currency,
		&i.MemberID,
		&i.CreatedAt,
		&i.StatusCode,
	)
	return i, err
}