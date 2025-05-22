package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kiriksik/TestTaskBrandScout/internal/config"
	"github.com/kiriksik/TestTaskBrandScout/internal/database"
	"github.com/kiriksik/TestTaskBrandScout/internal/model"
)

type QuotesService struct {
	ApiConfig *config.ApiConfig
}

func (quotesService *QuotesService) CreateQuote(ctx context.Context, req *model.QuoteRequest) (model.QuoteResponse, int, error) {
	if req == nil {
		return model.QuoteResponse{}, http.StatusBadRequest, fmt.Errorf("bad request")
	}

	quote, err := quotesService.ApiConfig.Queries.CreateQuote(ctx,
		database.CreateQuoteParams{
			Author: req.Author,
			Quote:  req.Quote,
		})
	if err != nil {
		return model.QuoteResponse{}, http.StatusInternalServerError, fmt.Errorf("error saving quote: %s", err)
	}

	fmt.Println("saved quote:", quote)

	return model.QuoteResponse(quote), http.StatusCreated, nil
}

func (quotesService *QuotesService) GetQuoteByID(ctx context.Context, id string) (model.QuoteResponse, int, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return model.QuoteResponse{}, 0, errors.New("invalid id format")
	}
	quote, err := quotesService.ApiConfig.Queries.GetQuoteByID(ctx, int32(idInt))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.QuoteResponse{}, http.StatusNotFound, fmt.Errorf("quote does not exists")
		}
		return model.QuoteResponse{}, http.StatusInternalServerError, fmt.Errorf("failed to get quote: %s", err)
	}
	fmt.Println("request for get quote:", quote)

	return model.QuoteResponse(quote), http.StatusOK, nil
}

func (quotesService *QuotesService) GetRandomQuote(ctx context.Context) (model.QuoteResponse, int, error) {
	quote, err := quotesService.ApiConfig.Queries.GetRandomQuote(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.QuoteResponse{}, http.StatusNotFound, fmt.Errorf("quotes does not exists")
		}
		return model.QuoteResponse{}, http.StatusInternalServerError, fmt.Errorf("failed to get quote: %s", err)
	}
	fmt.Println("request for get random quote:", quote)

	return model.QuoteResponse(quote), http.StatusOK, nil
}

func (quotesService *QuotesService) GetQuotes(ctx context.Context) ([]model.QuoteResponse, int, error) {

	quotes, err := quotesService.ApiConfig.Queries.GetQuotes(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []model.QuoteResponse{}, http.StatusNotFound, fmt.Errorf("quotes does not exists")
		}
		return []model.QuoteResponse{}, http.StatusInternalServerError, fmt.Errorf("failed to get quotes: %s", err)
	}
	fmt.Println("request for get quotes")

	convertedQuotes := make([]model.QuoteResponse, len(quotes))
	for i, quote := range quotes {
		convertedQuotes[i] = model.QuoteResponse(quote)
	}

	return convertedQuotes, http.StatusOK, nil
}

func (quotesService *QuotesService) GetQuotesByAuthor(ctx context.Context, author string) ([]model.QuoteResponse, int, error) {
	quotes, err := quotesService.ApiConfig.Queries.GetQuotesByAuthor(ctx, author)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to fetch quotes: %w", err)
	}

	convertedQuotes := make([]model.QuoteResponse, len(quotes))
	for i, quote := range quotes {
		convertedQuotes[i] = model.QuoteResponse(quote)
	}

	return convertedQuotes, http.StatusOK, nil
}

func (quotesService *QuotesService) DeleteQuote(ctx context.Context, id string) (model.QuoteResponse, int, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return model.QuoteResponse{}, 0, errors.New("invalid id format")
	}

	quote, err := quotesService.ApiConfig.Queries.DeleteQuote(ctx, int32(idInt))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.QuoteResponse{}, http.StatusNotFound, fmt.Errorf("quote does not exists")
		}
		return model.QuoteResponse{}, http.StatusInternalServerError, fmt.Errorf("failed to delete quote: %s", err)
	}
	fmt.Println("deleted quote:", quote)

	return model.QuoteResponse(quote), http.StatusOK, nil
}
