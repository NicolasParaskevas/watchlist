package data

import (
	"encoding/json"
	"log"
	"os"
)

type JSONSymbolRepository struct {
	filePath string
}

func NewJSONSymbolRepository(filePath string) *JSONSymbolRepository {
	return &JSONSymbolRepository{
		filePath: filePath,
	}
}
func (r *JSONSymbolRepository) GetAllSymbols() ([]Symbol, error) {
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
