-- name: GetMember :one
SELECT * FROM members
WHERE id = $1 
LIMIT 1;

-- name: GetMemberByEmail :one
SELECT * FROM members
WHERE email = $1 
LIMIT 1;

-- name: ListMembers :many
SELECT * FROM members
ORDER BY created_at;

-- name: ListMembersPaginated :many
SELECT * FROM members
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: CreateMember :one
INSERT INTO members (
  full_name, email, status_code
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateMember :one
UPDATE members 
SET full_name = $2, email = $3, status_code = $4
WHERE id=$1
RETURNING *;

-- name: UpdateMemberStaus :one
UPDATE members 
SET status_code = $2
WHERE id=$1
RETURNING *;

-- name: DeleteMember :exec
DELETE FROM members
WHERE id = $1;