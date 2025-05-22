package model

type QuoteRequest struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuoteResponse struct {
	ID     int32  `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}
