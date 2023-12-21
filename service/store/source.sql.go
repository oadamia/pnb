// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: source.sql

package store

import (
	"context"

	"github.com/lib/pq"
)

const createSource = `-- name: CreateSource :one
INSERT INTO source (id, name, url, description, category, language, country) 
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING name, url, created_at, updated_at, id, description, category, language, country
`

type CreateSourceParams struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

func (q *Queries) CreateSource(ctx context.Context, arg CreateSourceParams) (Source, error) {
	row := q.db.QueryRowContext(ctx, createSource,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.Description,
		arg.Category,
		arg.Language,
		arg.Country,
	)
	var i Source
	err := row.Scan(
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID,
		&i.Description,
		&i.Category,
		&i.Language,
		&i.Country,
	)
	return i, err
}

const deleteSource = `-- name: DeleteSource :one
DELETE FROM 
    source 
WHERE id = $1
RETURNING name, url, created_at, updated_at, id, description, category, language, country
`

func (q *Queries) DeleteSource(ctx context.Context, id string) (Source, error) {
	row := q.db.QueryRowContext(ctx, deleteSource, id)
	var i Source
	err := row.Scan(
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID,
		&i.Description,
		&i.Category,
		&i.Language,
		&i.Country,
	)
	return i, err
}

const getSource = `-- name: GetSource :one
SELECT name, url, created_at, updated_at, id, description, category, language, country 
FROM source 
WHERE id = $1
`

func (q *Queries) GetSource(ctx context.Context, id string) (Source, error) {
	row := q.db.QueryRowContext(ctx, getSource, id)
	var i Source
	err := row.Scan(
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID,
		&i.Description,
		&i.Category,
		&i.Language,
		&i.Country,
	)
	return i, err
}

const listSource = `-- name: ListSource :many
SELECT name, url, created_at, updated_at, id, description, category, language, country 
FROM source
WHERE 
    (NOT $1::boolean OR category = ANY($2::TEXT[]))
    AND (NOT $3::boolean OR country = ANY($4::TEXT[]))
    AND (NOT $5::boolean OR language = ANY($6::TEXT[]))
`

type ListSourceParams struct {
	CategoriesSet bool     `json:"categories_set"`
	Categories    []string `json:"categories"`
	CountriesSet  bool     `json:"countries_set"`
	Countries     []string `json:"countries"`
	LanguagesSet  bool     `json:"languages_set"`
	Languages     []string `json:"languages"`
}

func (q *Queries) ListSource(ctx context.Context, arg ListSourceParams) ([]Source, error) {
	rows, err := q.db.QueryContext(ctx, listSource,
		arg.CategoriesSet,
		pq.Array(arg.Categories),
		arg.CountriesSet,
		pq.Array(arg.Countries),
		arg.LanguagesSet,
		pq.Array(arg.Languages),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Source
	for rows.Next() {
		var i Source
		if err := rows.Scan(
			&i.Name,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID,
			&i.Description,
			&i.Category,
			&i.Language,
			&i.Country,
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

const updateSource = `-- name: UpdateSource :one
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
RETURNING name, url, created_at, updated_at, id, description, category, language, country
`

type UpdateSourceParams struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

func (q *Queries) UpdateSource(ctx context.Context, arg UpdateSourceParams) (Source, error) {
	row := q.db.QueryRowContext(ctx, updateSource,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.Description,
		arg.Category,
		arg.Language,
		arg.Country,
	)
	var i Source
	err := row.Scan(
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID,
		&i.Description,
		&i.Category,
		&i.Language,
		&i.Country,
	)
	return i, err
}
