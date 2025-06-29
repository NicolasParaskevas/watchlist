package data

import (
	"encoding/json"
	"log"
	"os"
)

type Symbol struct {
	Symbol string  `json:"symbol"`
	Type   string  `json:"type"`
	Name   string  `json:"name"`
	Price  float64 `json:"-"`
}

var cached []Symbol

func GetAllSymbols() ([]Symbol, error) {

	// INFO
	// Just using a variable for now to avoid reading
	// from the datasource each call since the list won't be
	// updated. This could be a REDIS cache in the future
	if cached != nil {
		return cached, nil
	}

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

	cached = symbols

	return symbols, nil
}
