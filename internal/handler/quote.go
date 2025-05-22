package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kiriksik/TestTaskBrandScout/internal/config"
	"github.com/kiriksik/TestTaskBrandScout/internal/model"
	"github.com/kiriksik/TestTaskBrandScout/internal/service"
)

type ApiHandler struct {
	ApiCfg *config.ApiConfig
}

type responseError struct {
	Err string `json:"error"`
}

func InitializeMux(ac *config.ApiConfig) *http.ServeMux {

	ah := &ApiHandler{ApiCfg: ac}
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("GET /quotes", ah.getQuotes)
	serveMux.HandleFunc("POST /quotes", ah.createQuote)
	serveMux.HandleFunc("GET /quotes/random", ah.getRandomQuote)
	serveMux.HandleFunc("DELETE /quotes/{quoteID}", ah.deleteQuote)
	return serveMux
}

func respondWithError(rw http.ResponseWriter, code int, errorMessage string) {

	rw.Header().Set("Content-Type", "application/json")

	rw.WriteHeader(code)

	responseError := responseError{Err: errorMessage}
	jsonErr, _ := json.Marshal(responseError)

	rw.Write(jsonErr)
}

func respondWithJson(rw http.ResponseWriter, code int, payload interface{}) {

	rw.Header().Set("Content-Type", "application/json")

	encodedJson, err := json.Marshal(payload)
	if err != nil {
		respondWithError(rw, http.StatusInternalServerError, fmt.Sprintf("error marshalling json: %s", err))
		return
	}

	rw.WriteHeader(code)
	rw.Write(encodedJson)
}

func (ah *ApiHandler) createQuote(rw http.ResponseWriter, req *http.Request) {

	quotesService := service.QuotesService{ApiConfig: ah.ApiCfg}
	var reqBodyData model.QuoteRequest

	err := json.NewDecoder(req.Body).Decode(&reqBodyData)
	defer req.Body.Close()
	if err != nil {
		respondWithError(rw, http.StatusBadRequest, fmt.Sprintf("error marshalling json: %s", err))
		return
	}

	quote, status, err := quotesService.CreateQuote(req.Context(), &reqBodyData)
	if err != nil {
		respondWithError(rw, status, err.Error())
		return
	}

	respondWithJson(rw, status, quote)
}

func (ah *ApiHandler) deleteQuote(rw http.ResponseWriter, req *http.Request) {
	quotesService := service.QuotesService{ApiConfig: ah.ApiCfg}
	quoteID := req.PathValue("quoteID")
	if quoteID == "" {
		respondWithError(rw, http.StatusBadRequest, "missing id")
		return
	}

	quote, status, err := quotesService.DeleteQuote(req.Context(), quoteID)
	if err != nil {
		respondWithError(rw, status, err.Error())
		return
	}

	respondWithJson(rw, status, quote)
}

func (ah *ApiHandler) getQuotes(rw http.ResponseWriter, req *http.Request) {
	quotesService := service.QuotesService{ApiConfig: ah.ApiCfg}

	author := req.URL.Query().Get("author")

	var (
		quotes interface{}
		status int
		err    error
	)

	if author != "" {
		quotes, status, err = quotesService.GetQuotesByAuthor(req.Context(), author)
	} else {
		quotes, status, err = quotesService.GetQuotes(req.Context())
	}
	if err != nil {
		respondWithError(rw, status, err.Error())
		return
	}

	respondWithJson(rw, status, quotes)
}

func (ah *ApiHandler) getRandomQuote(rw http.ResponseWriter, req *http.Request) {
	quotesService := service.QuotesService{ApiConfig: ah.ApiCfg}

	quote, status, err := quotesService.GetRandomQuote(req.Context())
	if err != nil {
		respondWithError(rw, status, err.Error())
		return
	}

	respondWithJson(rw, status, quote)
}
