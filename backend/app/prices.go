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
		log.Println("Error getting symbols")
		return
	}

	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		for _, symbol := range symbols {
			price := rand.Float64()*100 + 100 // Random price between 100-200
			h.BroadcastChan <- PriceUpdate{
				Symbol: symbol,
				Price:  price,
			}
		}
	}
}
