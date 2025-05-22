package service

import (
	"context"
	"testing"

	"github.com/kiriksik/TestTaskBrandScout/internal/config"
	"github.com/kiriksik/TestTaskBrandScout/internal/database"
	"github.com/kiriksik/TestTaskBrandScout/internal/model"
)

type mockQueries struct{}

func (m *mockQueries) CreateQuote(ctx context.Context, arg database.CreateQuoteParams) (database.Quote, error) {
	return database.Quote{ID: 1, Author: arg.Author, Quote: arg.Quote}, nil
}

func (m *mockQueries) GetQuoteByID(ctx context.Context, id int32) (database.Quote, error) {
	return database.Quote{ID: id, Author: "Test", Quote: "Test quote"}, nil
}

func (m *mockQueries) GetQuotes(ctx context.Context) ([]database.Quote, error) {
	return []database.Quote{
		{ID: 1, Author: "Author1", Quote: "Quote1"},
		{ID: 2, Author: "Author2", Quote: "Quote2"},
	}, nil
}

func (m *mockQueries) GetQuotesByAuthor(ctx context.Context, author string) ([]database.Quote, error) {
	if author == "Test" {
		return []database.Quote{
			{ID: 1, Author: "Test", Quote: "Test Quote"},
		}, nil
	}
	return []database.Quote{}, nil
}

func (m *mockQueries) DeleteQuote(ctx context.Context, id int32) (database.Quote, error) {
	return database.Quote{ID: id, Author: "Author", Quote: "Deleted Quote"}, nil
}

func (m *mockQueries) GetRandomQuote(ctx context.Context) (database.Quote, error) {
	return database.Quote{ID: 99, Author: "Random", Quote: "Random Quote"}, nil
}

func TestGetQuotes(t *testing.T) {
	service := &QuotesService{
		ApiConfig: &config.ApiConfig{
			Queries: &mockQueries{},
		},
	}

	quotes, status, err := service.GetQuotes(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Errorf("expected status 200, got %d", status)
	}
	if len(quotes) != 2 {
		t.Errorf("expected 2 quotes, got %d", len(quotes))
	}
}

func TestGetQuotesByAuthor(t *testing.T) {
	service := &QuotesService{
		ApiConfig: &config.ApiConfig{
			Queries: &mockQueries{},
		},
	}

	quotes, status, err := service.GetQuotesByAuthor(context.Background(), "Test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Errorf("expected status 200, got %d", status)
	}
	if len(quotes) != 1 || quotes[0].Author != "Test" {
		t.Errorf("unexpected quote result: %+v", quotes)
	}
}

func TestCreateQuote(t *testing.T) {
	service := &QuotesService{
		ApiConfig: &config.ApiConfig{
			Queries: &mockQueries{},
		},
	}

	req := &model.QuoteRequest{
		Author: "Tester",
		Quote:  "This is a test quote",
	}

	quote, status, err := service.CreateQuote(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 201 {
		t.Errorf("expected status 201, got %d", status)
	}
	if quote.Author != "Tester" {
		t.Errorf("unexpected author: %s", quote.Author)
	}
}
