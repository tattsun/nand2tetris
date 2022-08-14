package parser

import "io"

type CommandType string

type Parser struct {
}

func NewParser(r io.Reader) *Parser {
	return &Parser{}
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

func (p *Parser) Arg1() string {
	panic("not implemented")
}

func (p *Parser) Arg2() int64 {
	panic("not implemented")
}
