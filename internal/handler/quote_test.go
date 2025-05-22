package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kiriksik/TestTaskBrandScout/internal/config"
	"github.com/kiriksik/TestTaskBrandScout/internal/database"
)

type mockQueries struct{}

func (m *mockQueries) GetQuotes(ctx context.Context) ([]database.Quote, error) {
	return []database.Quote{
		{ID: 1, Author: "Author1", Quote: "Quote1"},
	}, nil
}

func (m *mockQueries) GetQuotesByAuthor(ctx context.Context, author string) ([]database.Quote, error) {
	return []database.Quote{
		{ID: 2, Author: author, Quote: "Quote by " + author},
	}, nil
}

func (m *mockQueries) CreateQuote(ctx context.Context, arg database.CreateQuoteParams) (database.Quote, error) {
	return database.Quote{ID: 3, Author: arg.Author, Quote: arg.Quote}, nil
}

func (m *mockQueries) GetRandomQuote(ctx context.Context) (database.Quote, error) {
	return database.Quote{ID: 99, Author: "Random", Quote: "Random quote"}, nil
}

func (m *mockQueries) DeleteQuote(ctx context.Context, id int32) (database.Quote, error) {
	return database.Quote{ID: id, Author: "Deleted", Quote: "Deleted quote"}, nil
}

func (m *mockQueries) GetQuoteByID(ctx context.Context, id int32) (database.Quote, error) {
	return database.Quote{ID: id, Author: "Author", Quote: "Quote"}, nil
}

func TestGetQuotesHandler(t *testing.T) {
	apiCfg := &config.ApiConfig{
		Queries: &mockQueries{},
	}
	handler := &ApiHandler{ApiCfg: apiCfg}

	req := httptest.NewRequest(http.MethodGet, "/quotes", nil)
	rr := httptest.NewRecorder()

	handler.getQuotes(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
	expected := `"Author1"`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("expected body to contain %s, got %s", expected, rr.Body.String())
	}
}

func TestCreateQuoteHandler(t *testing.T) {
	apiCfg := &config.ApiConfig{
		Queries: &mockQueries{},
	}
	handler := &ApiHandler{ApiCfg: apiCfg}

	body := `{"author": "Tester", "quote": "This is a test"}`
	req := httptest.NewRequest(http.MethodPost, "/quotes", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.createQuote(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rr.Code)
	}

	expected := `"Tester"`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("expected body to contain %s, got %s", expected, rr.Body.String())
	}
}
