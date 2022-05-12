-- name: GetDeposit :one
SELECT * FROM deposits
WHERE id = $1 
LIMIT 1;

-- name: GetDepositByReferenceNum :one
SELECT * FROM deposits
WHERE reference_num = $1 
LIMIT 1;

-- name: ListDeposits :many
SELECT * FROM deposits
ORDER BY created_at;

-- name: ListDepositsByMember :many
SELECT * FROM deposits
WHERE member_id = $1
ORDER BY created_at;

-- name: ListDepositsPaginated :many
SELECT * FROM deposits
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: CreateDeposit :one
INSERT INTO deposits (
  reference_num, amount, currency, member_id, status_code
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateDeposit :one
UPDATE deposits 
SET reference_num = $2, amount = $3, currency = $4, member_id = $5, status_code = $6
WHERE id=$1
RETURNING *;

-- name: UpdateDepositStaus :one
UPDATE deposits 
SET status_code = $2
WHERE id=$1
RETURNING *;

-- name: DeleteDeposit :exec
DELETE FROM deposits
WHERE id = $1;