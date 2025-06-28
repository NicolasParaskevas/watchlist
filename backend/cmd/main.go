package main

import (
	"log"

	"github.com/nicolasparaskevas/watchlist/app"
)

func main() {

	server := app.NewServer()
	if err := server.Start(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
