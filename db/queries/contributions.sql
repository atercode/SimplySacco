CREATE TABLE "contributions" (
  "id" SERIAL PRIMARY KEY,
  "deposit_id" int,
  "amount" bigint NOT NULL,
  "currency" varchar DEFAULT 'KES',
  "member_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "status_code" varchar NOT NULL
);

-- name: GetContribution :one
SELECT * FROM contributions
WHERE id = $1 
LIMIT 1;

-- name: ListContributions :many
SELECT * FROM contributions
ORDER BY created_at;

-- name: ListContributionsByMember :many
SELECT * FROM contributions
WHERE member_id = $1
ORDER BY created_at;

-- name: ListContributionsByDeposit :many
SELECT * FROM contributions
WHERE deposit_id = $1
ORDER BY created_at;

-- name: ListContributionsPaginated :many
SELECT * FROM contributions
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: CreateContribution :one
INSERT INTO contributions (
  deposit_id, amount, currency, member_id, status_code
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateContribution :one
UPDATE contributions 
SET deposit_id = $2, amount = $3, currency = $4, member_id = $5, status_code = $6
WHERE id=$1
RETURNING *;

-- name: UpdateContributionStaus :one
UPDATE contributions 
SET status_code = $2
WHERE id=$1
RETURNING *;

-- name: DeleteContribution :exec
DELETE FROM contributions
WHERE id = $1;