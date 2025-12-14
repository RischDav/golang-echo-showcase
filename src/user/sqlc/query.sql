-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;

-- name: ListUser :many
SELECT * FROM user
ORDER BY name;

-- name: CreateUser :one
INSERT INTO user (
  firstname, lastname
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateUserFirstname :exec
UPDATE user
set firstname = ?
WHERE id = ?;

-- name: UpdateUserLastname :exec
UPDATE user
set lastname = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;