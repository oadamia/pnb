-- +goose Up
-- +goose StatementBegin
ALTER TABLE sources
DROP CONSTRAINT sources_pkey,
DROP COLUMN id,
DROP COLUMN driver,
ADD COLUMN id TEXT PRIMARY KEY NOT NULL,
ADD COLUMN description TEXT NOT NULL,
ADD COLUMN category TEXT NOT NULL,
ADD COLUMN language CHAR(2) NOT NULL,
ADD COLUMN country CHAR(2) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
