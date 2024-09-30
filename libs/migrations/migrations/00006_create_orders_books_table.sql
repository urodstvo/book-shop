-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS orders_books (
    order_id INTEGER NOT NULL,
    book_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    price NUMERIC NOT NULL,

    PRIMARY KEY (order_id, book_id),
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS orders_books;
-- +goose StatementEnd
