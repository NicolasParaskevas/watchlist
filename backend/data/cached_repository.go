// data/symbol_cache_repo.go
package data

import (
	"log"
	"sync"
)

type CachedSymbolRepository struct {
	inner SymbolRepository
	cache []Symbol
	mu    sync.RWMutex
}

func NewCachedSymbolRepository(inner SymbolRepository) *CachedSymbolRepository {
	return &CachedSymbolRepository{
		inner: inner,
	}
}

func (r *CachedSymbolRepository) GetAllSymbols() ([]Symbol, error) {
	// INFO
	// Just using a variable for now to avoid reading
	// from the datasource each call since the list won't be
	// updated. This could be a REDIS cache in the future
	r.mu.RLock()
	if r.cache != nil {
		defer r.mu.RUnlock()
		log.Println("getting from cache!")
		return r.cache, nil
	}
	r.mu.RUnlock()

	r.mu.Lock()
	defer r.mu.Unlock()

	symbols, err := r.inner.GetAllSymbols()
	if err != nil {
		return nil, err
	}

	r.cache = symbols
	return symbols, nil
}
