-- name: GetStatus :one
SELECT * FROM statuses
WHERE code = $1 LIMIT 1;

-- name: ListStatuses :many
SELECT * FROM statuses
ORDER BY name;

-- name: ListStatusesPaginated :many
SELECT * FROM statuses
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: CreateStatus :one
INSERT INTO statuses (
  code, name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateStatus :one
UPDATE statuses 
SET name = $2 
WHERE code=$1
RETURNING *;

-- name: DeleteStatus :exec
DELETE FROM statuses
WHERE code = $1;