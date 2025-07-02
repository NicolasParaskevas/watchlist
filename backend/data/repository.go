package data

type SymbolRepository interface {
	GetAllSymbols() ([]Symbol, error)
}
