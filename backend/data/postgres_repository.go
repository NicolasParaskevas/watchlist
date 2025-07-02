package data

import (
	"database/sql"
)

type PostgresSymbolRepository struct {
	db *sql.DB
}

func NewPostgresSymbolRepository(db *sql.DB) *PostgresSymbolRepository {
	return &PostgresSymbolRepository{db: db}
}

func (r *PostgresSymbolRepository) GetAllSymbols() ([]Symbol, error) {
	// TODO: Implement actual DB fetching
	return nil, nil
}
