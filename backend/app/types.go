package app

type PriceUpdate struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}
