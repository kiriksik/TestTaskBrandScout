package database

import (
	"context"
)

type QuotesQueries interface {
	CreateQuote(ctx context.Context, arg CreateQuoteParams) (Quote, error)
	GetQuoteByID(ctx context.Context, id int32) (Quote, error)
	GetQuotes(ctx context.Context) ([]Quote, error)
	GetQuotesByAuthor(ctx context.Context, author string) ([]Quote, error)
	DeleteQuote(ctx context.Context, id int32) (Quote, error)
	GetRandomQuote(ctx context.Context) (Quote, error)
}
