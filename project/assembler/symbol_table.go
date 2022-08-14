package main

type SymbolTable map[string]int64

func NewSymbolTable() SymbolTable {
	return make(SymbolTable)
}

func (t SymbolTable) AddEntry(symbol string, address int64) {
	t[symbol] = address
}

func (t SymbolTable) Contains(symbol string) bool {
	_, ok := t[symbol]
	return ok
}

func (t SymbolTable) GetAddress(symbol string) int64 {
	return t[symbol]
}
