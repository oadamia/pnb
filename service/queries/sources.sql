-- name: InsertSource :one
INSERT INTO sources (name, url, driver) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListSources :many
SELECT * 
FROM sources;

-- name: SelectSource :one
SELECT * 
FROM sources 
WHERE id = $1;

-- name: UpdateSource :one
UPDATE sources 
SET 
    name = $2, 
    url = $3, 
    driver = $4,
    created_at = $5
WHERE id = $1
RETURNING *;

-- name: DeleteSource :one
DELETE FROM 
    sources 
WHERE id = $1
RETURNING *;