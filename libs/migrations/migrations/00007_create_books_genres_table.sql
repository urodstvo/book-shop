-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS books_genres (
    book_id INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,

    PRIMARY KEY (genre_id, book_id),
    FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS books_genres;
-- +goose StatementEnd
