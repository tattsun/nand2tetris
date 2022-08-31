package tokenizer

import "io"

type Tokenizer struct {
	r io.Reader
}

func NewTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{
		r: r,
	}
}

func (t *Tokenizer) HasMoreTokens() bool {
	panic("not implemented")
}

func (t *Tokenizer) Advance() {
	panic("not implemented")
}

func (t *Tokenizer) TokenType() string {
	panic("not implemented")
}

func (t *Tokenizer) KeyWord() string {
	panic("not implemented")
}

func (t *Tokenizer) Symbol() rune {
	panic("not implemented")
}

func (t *Tokenizer) Identifier() string {
	panic("not implemented")
}

func (t *Tokenizer) IntVal() int {
	panic("not implemented")
}

func (t *Tokenizer) StringVal() string {
	panic("not implemented")
}
