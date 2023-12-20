-- name: CreateSource :one
INSERT INTO source (id, name, url, description, category, language, country) 
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: ListSources :many
SELECT * 
FROM source
WHERE 
    (NOT @categories_set::boolean OR category = ANY(@categories::TEXT[]))
    AND (NOT @countries_set::boolean OR country = ANY(@countries::TEXT[]));

-- name: GetSource :one
SELECT * 
FROM source 
WHERE id = $1;

-- name: UpdateSource :one
UPDATE source 
SET 
    name = $2, 
    url = $3, 
    description = $4,
    category = $5,
    language = $6,
    country = $7,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteSource :one
DELETE FROM 
    source 
WHERE id = $1
RETURNING *;
