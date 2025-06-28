package app

import (
	"encoding/json"
	"log"
	"os"
)

type Symbol struct {
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
	Name   string `json:"name"`
}

func GetAllSymbols() ([]Symbol, error) {
	// TODO use DB later
	data, err := os.ReadFile("data/symbols.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var symbols []Symbol
	if err := json.Unmarshal(data, &symbols); err != nil {
		log.Println(err)
		return nil, err
	}

	return symbols, nil
}
