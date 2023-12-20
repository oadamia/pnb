-- +goose Up
-- +goose StatementBegin
ALTER TABLE sources
Rename TO source;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
