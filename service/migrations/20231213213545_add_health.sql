-- +goose Up
-- +goose StatementBegin
CREATE TABLE health
(
    message TEXT NOT NULL,
    service TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS health;
-- +goose StatementEnd
