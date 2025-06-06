// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: quotes.sql

package database

import (
	"context"
)

const createQuote = `-- name: CreateQuote :one
INSERT INTO quotes (author, quote)
VALUES (
    $1,
    $2
) RETURNING id, author, quote
`

type CreateQuoteParams struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func (q *Queries) CreateQuote(ctx context.Context, arg CreateQuoteParams) (Quote, error) {
	row := q.db.QueryRowContext(ctx, createQuote, arg.Author, arg.Quote)
	var i Quote
	err := row.Scan(&i.ID, &i.Author, &i.Quote)
	return i, err
}

const deleteQuote = `-- name: DeleteQuote :one
DELETE FROM quotes
WHERE id = $1
RETURNING id, author, quote
`

func (q *Queries) DeleteQuote(ctx context.Context, id int32) (Quote, error) {
	row := q.db.QueryRowContext(ctx, deleteQuote, id)
	var i Quote
	err := row.Scan(&i.ID, &i.Author, &i.Quote)
	return i, err
}

const getQuoteByID = `-- name: GetQuoteByID :one
SELECT id, author, quote FROM quotes
WHERE id = $1
`

func (q *Queries) GetQuoteByID(ctx context.Context, id int32) (Quote, error) {
	row := q.db.QueryRowContext(ctx, getQuoteByID, id)
	var i Quote
	err := row.Scan(&i.ID, &i.Author, &i.Quote)
	return i, err
}

const getQuotes = `-- name: GetQuotes :many
SELECT id, author, quote FROM quotes
`

func (q *Queries) GetQuotes(ctx context.Context) ([]Quote, error) {
	rows, err := q.db.QueryContext(ctx, getQuotes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quote
	for rows.Next() {
		var i Quote
		if err := rows.Scan(&i.ID, &i.Author, &i.Quote); err != nil {
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

const getQuotesByAuthor = `-- name: GetQuotesByAuthor :many
SELECT id, author, quote FROM quotes
WHERE author = $1
`

func (q *Queries) GetQuotesByAuthor(ctx context.Context, author string) ([]Quote, error) {
	rows, err := q.db.QueryContext(ctx, getQuotesByAuthor, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quote
	for rows.Next() {
		var i Quote
		if err := rows.Scan(&i.ID, &i.Author, &i.Quote); err != nil {
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

const getRandomQuote = `-- name: GetRandomQuote :one
SELECT id, author, quote FROM quotes 
ORDER BY RANDOM() 
LIMIT 1
`

func (q *Queries) GetRandomQuote(ctx context.Context) (Quote, error) {
	row := q.db.QueryRowContext(ctx, getRandomQuote)
	var i Quote
	err := row.Scan(&i.ID, &i.Author, &i.Quote)
	return i, err
}
