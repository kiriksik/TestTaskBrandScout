-- +goose Up
CREATE TABLE IF NOT EXISTS quotes (
    id SERIAL PRIMARY KEY,
    author TEXT NOT NULL,
    quote TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS quotes;