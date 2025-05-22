-- name: CreateQuote :one
INSERT INTO quotes (author, quote)
VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetQuotes :many
SELECT * FROM quotes;

-- name: GetQuotesByAuthor :many
SELECT * FROM quotes
WHERE author = $1;


-- name: GetQuoteByID :one
SELECT * FROM quotes
WHERE id = $1;


-- name: DeleteQuote :one
DELETE FROM quotes
WHERE id = $1
RETURNING *;

-- name: GetRandomQuote :one
SELECT * FROM quotes 
ORDER BY RANDOM() 
LIMIT 1;
