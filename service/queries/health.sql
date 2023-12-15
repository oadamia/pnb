-- name: CreateHealth :one
INSERT INTO health (
    message, 
    service
) 
VALUES (
    $1, 
    $2
)
RETURNING *;

