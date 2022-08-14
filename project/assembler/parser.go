package main

import (
	"fmt"
	"io"
)

type CommandType string

const (
	INVALID_COMMAND CommandType = "INVALID_COMMAND"
	A_COMMAND       CommandType = "A_COMMAND"
	C_COMMAND       CommandType = "C_COMMAND"
	L_COMMAND       CommandType = "L_COMMAND"
)

type token struct {
	CommandType CommandType
	Symbol      string
	Dest        string
	Comp        string
	Jump        string
}

func (t *token) String() string {
	if t.CommandType == A_COMMAND {
		return fmt.Sprintf("@%s", t.Symbol)
	} else if t.CommandType == C_COMMAND {
		str := ""
		if len(t.Dest) > 0 {
			str += t.Dest + "="
		}
		str += t.Comp
		if len(t.Jump) > 0 {
			str += ";" + t.Jump
		}
		return str
	} else {
		return fmt.Sprintf("(%s)", t.Symbol)
	}

}

type Parser struct {
	tokenizer    *LookAheadTokenizer
	currentToken *token
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		tokenizer: NewLookAheadTokenizer(NewTokenizer(r)),
	}
}

func (p *Parser) HasMoreCommands() bool {
	_, err := p.tokenizer.LookAhead()
	return err != io.EOF
}

func (p *Parser) Advance() error {
	lookAhead, err := p.tokenizer.LookAhead()
	if err != nil {
		return err
	}

	for lookAhead.Type == TOKEN_TYPE_LINEBREAK {
		_, err := p.mustToken(TOKEN_TYPE_LINEBREAK)
		if err != nil {
			return err
		}
		continue
	}

	if lookAhead.Type == TOKEN_TYPE_ATSIGN {
		p.currentToken, err = p.mustACommand()
		return err
	}

	if lookAhead.Type == TOKEN_TYPE_SYMBOL {
		p.currentToken, err = p.mustCCommand()
		return err
	}

	if lookAhead.Type == TOKEN_TYPE_L_PAREN {
		p.currentToken, err = p.mustLCommand()
		return err
	}

	return fmt.Errorf("unexpected token: %+v", lookAhead)
}

func (p *Parser) mustCCommand() (*token, error) {
	var snd string

	t := &token{
		CommandType: C_COMMAND,
	}

	fst, err := p.mustComp()
	if err != nil {
		return nil, err
	}

	lookAhead, _ := p.tokenizer.LookAhead()
	if lookAhead != nil && lookAhead.Type != TOKEN_TYPE_LINEBREAK && lookAhead.Type == TOKEN_TYPE_EQUAL {
		t.Dest = fst

		if _, err := p.mustToken(TOKEN_TYPE_EQUAL); err != nil {
			return nil, err
		}

		snd, err = p.mustComp()
		if err != nil {
			return nil, err
		}
		t.Comp = snd
	} else {
		t.Comp = fst
	}

	lookAhead, _ = p.tokenizer.LookAhead()
	if lookAhead != nil && lookAhead.Type != TOKEN_TYPE_LINEBREAK && lookAhead.Type == TOKEN_TYPE_SEMICOLON {
		if _, err := p.mustToken(TOKEN_TYPE_SEMICOLON); err != nil {
			return nil, err
		}

		jmp, err := p.mustToken(TOKEN_TYPE_SYMBOL)
		if err != nil {
			return nil, err
		}
		t.Jump = jmp
	}

	return t, nil
}

func (p *Parser) mustLCommand() (*token, error) {
	_, err := p.mustToken(TOKEN_TYPE_L_PAREN)
	if err != nil {
		return nil, err
	}

	sym, err := p.mustToken(TOKEN_TYPE_SYMBOL)
	if err != nil {
		return nil, err
	}

	_, err = p.mustToken(TOKEN_TYPE_R_PAREN)
	if err != nil {
		return nil, err
	}

	return &token{
		CommandType: L_COMMAND,
		Symbol:      sym,
	}, nil
}

func (p *Parser) mustACommand() (*token, error) {
	_, err := p.mustToken(TOKEN_TYPE_ATSIGN)
	if err != nil {
		return nil, err
	}

	sym, err := p.mustToken(TOKEN_TYPE_SYMBOL)
	if err != nil {
		return nil, err
	}

	return &token{
		CommandType: A_COMMAND,
		Symbol:      sym,
	}, nil
}

func (p *Parser) mustComp() (string, error) {
	lookAhead, _ := p.tokenizer.LookAhead()
	if lookAhead == nil || lookAhead.Type == TOKEN_TYPE_LINEBREAK {
		return "", fmt.Errorf("expected comp but got %+v", lookAhead)
	}

	if lookAhead.Type == TOKEN_TYPE_MINUS {
		if _, err := p.mustToken(TOKEN_TYPE_MINUS); err != nil {
			return "", err
		}

		reg, err := p.mustToken(TOKEN_TYPE_SYMBOL)
		if err != nil {
			return "", err
		}

		return "-" + reg, nil
	}

	if lookAhead.Type == TOKEN_TYPE_REVERSE {
		if _, err := p.mustToken(TOKEN_TYPE_REVERSE); err != nil {
			return "", err
		}

		reg, err := p.mustToken(TOKEN_TYPE_SYMBOL)
		if err != nil {
			return "", err
		}

		return "!" + reg, nil
	}

	if lookAhead.Type == TOKEN_TYPE_SYMBOL {
		left, err := p.mustToken(TOKEN_TYPE_SYMBOL)
		if err != nil {
			return "", err
		}

		lookAhead, _ := p.tokenizer.LookAhead()
		if lookAhead == nil || !lookAhead.IsOperator() {
			return left, nil
		}

		operator, err := p.mustOperator()
		if err != nil {
			return "", err
		}

		right, err := p.mustToken(TOKEN_TYPE_SYMBOL)
		if err != nil {
			return "", err
		}

		return left + operator + right, nil
	}

	return "", nil
}

func (p *Parser) mustJump() (string, error) {
	return p.mustToken(TOKEN_TYPE_SYMBOL)
}

func (p *Parser) mustOperator() (string, error) {
	lookAhead, err := p.tokenizer.LookAhead()
	if err != nil {
		return "", err
	}

	if lookAhead.IsOperator() {
		_, err := p.tokenizer.Read()
		return lookAhead.Value, err
	}

	return "", fmt.Errorf("expected operator but got: %s", lookAhead.Value)
}

func (p *Parser) mustToken(typ TokenType) (string, error) {
	t, err := p.tokenizer.Read()
	if err != nil {
		return "", err
	}

	if t.Type != typ {
		return "", fmt.Errorf("expected a %s but got: %+v", typ, t)
	}

	return t.Value, nil
}

func (p *Parser) CommandType() CommandType {
	return p.currentToken.CommandType
}

func (p *Parser) Symbol() string {
	return p.currentToken.Symbol
}

func (p *Parser) Dest() string {
	return p.currentToken.Dest
}

func (p *Parser) Comp() string {
	return p.currentToken.Comp
}

func (p *Parser) Jump() string {
	return p.currentToken.Jump
}
