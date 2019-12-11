package service

import (
	"context"

	api "github.com/atreyagaurav/quotes/pkg/api"
)

type quotesServer struct {
	name string
}

func NewQuotesServer(st *string) api.QuotesServer {
	return &quotesServer{name: *st}
}

func (q *quotesServer) Get(ctx context.Context, req *api.QuoteRequest) (*api.QuoteResponse, error) {
	return &api.QuoteResponse{
		Id:     1,
		Genre:  "test",
		Body:   "This is a test",
		Author: "gaurav",
	}, nil
}
