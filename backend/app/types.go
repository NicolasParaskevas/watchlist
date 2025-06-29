package app

import "github.com/nicolasparaskevas/watchlist/data"

type PriceUpdate struct {
	Symbol data.Symbol `json:"symbol"`
	Price  float64     `json:"price"`
}

type SubscribeMessage struct {
	Symbol   string `json:"symbol"`
	ClientId string `json:"client_id"`
}

type UnsubscribeMessage struct {
	Symbol   string `json:"symbol"`
	ClientId string `json:"client_id"`
}

type ClientMessage struct {
	Action   string `json:"action"` // "subscribe" or "unsubscribe"
	Symbol   string `json:"symbol"`
	ClientID string `json:"clientId"`
}
