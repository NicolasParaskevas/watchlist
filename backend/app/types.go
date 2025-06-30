package app

type PriceUpdate struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
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
