package main

import "io"

type CommandType string

const (
	A_COMMAND CommandType = "A_COMMAND"
	C_COMMAND CommandType = "C_COMMAND"
	L_COMMAND CommandType = "L_COMMAND"
)

type Parser struct {
	r io.Reader
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		r: r,
	}
}

func (p *Parser) HasMoreCommands() bool {
	panic("not implemented")
}

func (p *Parser) Advance() error {
	panic("not implemented")
}

func (p *Parser) CommandType() CommandType {
	panic("not implemented")
}

func (p *Parser) Symbol() string {
	panic("not implemented")
}

func (p *Parser) Dest() string {
	panic("not implemented")
}

func (p *Parser) Comp() string {
	panic("not implemented")
}

func (p *Parser) Jump() string {
	panic("not implemented")
}
