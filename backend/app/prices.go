package app

import (
	"log"
	"math/rand"
	"time"

	"github.com/nicolasparaskevas/watchlist/data"
)

func GetPriceData(h *Hub) {
	symbols, err := data.GetAllSymbols()

	if err != nil {
		log.Printf("Error getting symbols: %v", err)
		return
	}

	// for storing the previous prices, this helps with
	// simulating a more natural price fluctioation

	prices := make(map[string]float64)

	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		for _, symbol := range symbols {
			last := prices[symbol.Symbol]
			newPrice := generateRandomPrice(last)
			prices[symbol.Symbol] = newPrice

			h.BroadcastChan <- PriceUpdate{
				Symbol: symbol,
				Price:  newPrice,
			}
		}
	}
}

func generateRandomPrice(last float64) float64 {
	if last == 0 {
		// we set a random price between 100 and 200
		return rand.Float64()*100 + 100
	}

	// if the price is not 0 then we can use it and flucuate
	// some decimal points either up or down as to not send
	// a totally random number each time and simulate a real market
	// price fluctuation of +/- 0.5 percent

	changePercent := (rand.Float64() - 0.5) * 0.01
	delta := last * changePercent
	// round to 4 decimals for the "forex" look
	return float64(int((last+delta)*10000)) / 10000
}
