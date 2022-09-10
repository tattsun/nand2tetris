package compilationengine

import "fmt"

func (e *CompilationEngine) advance() {
	if !e.t.HasMoreTokens() {
		panic("expected token, but EOF")
	}
	e.t.Advance()
}

func (e *CompilationEngine) mustKeyword(keyword string) {
	e.advance()

	if e.t.TokenType() != "KEYWORD" {
		panic(fmt.Sprintf("expected keyword %s, but got token %s", keyword, e.t.TokenType()))
	}

	if e.t.KeyWord() != keyword {
		panic(fmt.Sprintf("expected keyword %s, but got keyword %s", keyword, e.t.KeyWord()))
	}
}

func (e *CompilationEngine) mustIdentifier() *Identifier {
	e.advance()

	if e.t.TokenType() != "IDENTIFIER" {
		panic(fmt.Sprintf("expected identifier, but got token %s", e.t.TokenType()))
	}

	return &Identifier{
		Value: e.t.Identifier(),
	}
}

func (e *CompilationEngine) mustSymbol(symbol byte) {
	e.advance()

	if e.t.TokenType() != "SYMBOL" {
		panic(fmt.Sprintf("expected symbol %c, but got token %s", symbol, e.t.TokenType()))
	}

	if e.t.RawSymbol() != symbol {
		panic(fmt.Sprintf("expected keyword %c, but got keyword %c", symbol, e.t.RawSymbol()))
	}
}
