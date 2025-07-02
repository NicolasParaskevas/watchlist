package main

import (
	"log"

	"github.com/nicolasparaskevas/watchlist/app"
	"github.com/nicolasparaskevas/watchlist/data"
)

func main() {

	jsonRepo := data.NewJSONSymbolRepository("data/symbols.json")
	cachedRepo := data.NewCachedSymbolRepository(jsonRepo)

	hub := app.NewHub()

	server := app.NewServer(hub, cachedRepo)

	if err := server.Start(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
