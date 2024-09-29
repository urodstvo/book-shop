-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS genres (
     id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
     name VARCHAR(255) NOT NULL UNIQUE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS genres;
-- +goose StatementEnd
