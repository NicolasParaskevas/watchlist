package app

import "github.com/nicolasparaskevas/watchlist/data"

type PriceUpdate struct {
	Symbol data.Symbol `json:"symbol"`
	Price  float64     `json:"price"`
}
