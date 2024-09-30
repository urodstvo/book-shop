-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS books (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL,
    cover TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    rating NUMERIC DEFAULT NULL,
    rating_count INTEGER DEFAULT 0,
    annotation TEXT NOT NULL,
    page_count INTEGER NOT NULL,
    stock_count INTEGER NOT NULL,
    orders_count INTEGER NOT NULL,

    published_by TEXT NOT NULL,
    published_at DATETIME NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
