package main

import (
	"fmt"
	"io"
	"unicode"
)

type TokenType string

const (
	TOKEN_TYPE_SYMBOL    TokenType = "TOKEN_TYPE_SYMBOL"    // hoge
	TOKEN_TYPE_ATSIGN    TokenType = "TOKEN_TYPE_ATSIGN"    // @
	TOKEN_TYPE_L_PAEN    TokenType = "TOKEN_TYPE_L_PAREN"   // (
	TOKEN_TYPE_R_PAEN    TokenType = "TOKEN_TYPE_R_PAREN"   // )
	TOKEN_TYPE_PLUS      TokenType = "TOKEN_TYPE_PLUS"      // +
	TOKEN_TYPE_MINUS     TokenType = "TOKEN_TYPE_MINUS"     // -
	TOKEN_TYPE_EQUAL     TokenType = "TOKEN_TYPE_EQUAL"     // =
	TOKEN_TYPE_REVERSE   TokenType = "TOKEN_TYPE_REVERSE"   // !
	TOKEN_TYPE_AND       TokenType = "TOKEN_TYPE_AND"       // &
	TOKEN_TYPE_OR        TokenType = "TOKEN_TYPE_OR"        // |
	TOKEN_TYPE_SEMICOLON TokenType = "TOKEN_TYPE_SEMICOLON" // ;
	TOKEN_TYPE_LINEBREAK TokenType = "TOKEN_TYPE_LINEBREAK" // \n
)

type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) IsOperator() bool {
	return t.Type == TOKEN_TYPE_PLUS || t.Type == TOKEN_TYPE_MINUS ||
		t.Type == TOKEN_TYPE_AND || t.Type == TOKEN_TYPE_OR
}

type Tokenizer struct {
	r         io.Reader
	lookAhead *rune
}

func NewTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{r: r}
}

func (t *Tokenizer) readRune() (rune, error) {
	if t.lookAhead != nil {
		r := *t.lookAhead
		t.lookAhead = nil
		return r, nil
	}

	buf := make([]byte, 1)

	_, err := t.r.Read(buf)
	if err != nil {
		return ' ', err
	}

	return rune(buf[0]), nil
}

func (t *Tokenizer) putBackRune(r rune) {
	t.lookAhead = &r
}

func (t *Tokenizer) Read() (Token, error) {
	token, err := t.read()
	return token, err
}

func (t *Tokenizer) read() (Token, error) {
	for {
		r, err := t.readRune()
		if err != nil {
			return Token{}, err
		}

		if unicode.IsSpace(r) {
			continue
		}

		if r == '\n' {
			return Token{
				Type:  TOKEN_TYPE_LINEBREAK,
				Value: "\n",
			}, nil
		}

		if r == '/' {
			r, err := t.readRune()
			if err != nil {
				return Token{}, err
			}

			if r != '/' {
				return Token{}, fmt.Errorf("expected // but got /")
			}

			for r != '\n' {
				r, err = t.readRune()
				if err != nil {
					return Token{}, err
				}
			}

			continue
		}

		if r == '@' {
			return Token{
				Type:  TOKEN_TYPE_ATSIGN,
				Value: "@",
			}, nil
		}

		if r == '(' {
			return Token{
				Type:  TOKEN_TYPE_L_PAEN,
				Value: "(",
			}, nil
		}

		if r == ')' {
			return Token{
				Type:  TOKEN_TYPE_R_PAEN,
				Value: ")",
			}, nil
		}

		if r == '+' {
			return Token{
				Type:  TOKEN_TYPE_PLUS,
				Value: "+",
			}, nil
		}

		if r == '-' {
			return Token{
				Type:  TOKEN_TYPE_MINUS,
				Value: "-",
			}, nil
		}

		if r == '=' {
			return Token{
				Type:  TOKEN_TYPE_EQUAL,
				Value: "=",
			}, nil
		}

		if r == '!' {
			return Token{
				Type:  TOKEN_TYPE_REVERSE,
				Value: "!",
			}, nil
		}

		if r == '&' {
			return Token{
				Type:  TOKEN_TYPE_AND,
				Value: "&",
			}, nil
		}

		if r == '|' {
			return Token{
				Type:  TOKEN_TYPE_OR,
				Value: "|",
			}, nil
		}

		if r == ';' {
			return Token{
				Type:  TOKEN_TYPE_SEMICOLON,
				Value: ";",
			}, nil
		}

		if IsAlpha(r) || IsSymbol(r) || IsNumber(r) {
			sym := string(r)

			for {
				r, err := t.readRune()
				if err != nil {
					if err == io.EOF {
						break
					}
					return Token{}, err
				}

				if IsAlpha(r) || IsNumber(r) || IsSymbol(r) {
					sym += string(r)
					continue
				}

				t.putBackRune(r)
				break
			}

			return Token{
				Type:  TOKEN_TYPE_SYMBOL,
				Value: sym,
			}, nil
		}

		return Token{}, fmt.Errorf("unexpected character: %s", string(r))
	}
}
