package main

import (
	"context"
	"fmt"
)

// PriceFetcher is an interface that can fetch a price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the price fetcher interface
type priceFetcher struct{}

// Business logic clean + only handles business logic

func (p *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker) // context allows us to cancel if the request takes too long
}

var priceMocks = map[string]float64{
	"RTC": 20_000.0,
	"HTH": 10_000.0,
	"GG":  100_000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]

	if !ok {
		return price, fmt.Errorf("The given ticker (%s) is not supported", ticker)
	}

	return price, nil
}
