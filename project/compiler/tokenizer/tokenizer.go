package tokenizer

import (
	"compiler/utils"
	"fmt"
	"io"
	"strconv"
)

type tokenInfo struct {
	TokenType  string
	Identifier string
	Keyword    string
	Symbol     byte
	IntVal     int
	StringVal  string
}

type Tokenizer struct {
	r            *LookaheadReader
	nextToken    *tokenInfo
	currentToken *tokenInfo
}

func NewTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{
		r: NewLookaheadReader(r),
	}
}

func (t *Tokenizer) HasMoreTokens() bool {
	if t.nextToken == nil {
		t.nextToken = t.readToken()
	}

	return t.nextToken != nil
}

func (t *Tokenizer) Advance() {
	if t.nextToken == nil {
		t.nextToken = t.readToken()
	}

	t.currentToken = t.nextToken
	t.nextToken = nil
}

func (t *Tokenizer) readToken() *tokenInfo {
	token := ""
	buf := make([]byte, 1)

	for {
		_, err := t.r.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				panic(err)
			}
		}

		// skip whitespaces
		if buf[0] == ' ' || buf[0] == '\t' || buf[0] == '\n' {
			continue
		}

		// skip comments
		if buf[0] == '/' {
			la, ok := t.r.Lookahead()
			if ok && (la == '/' || la == '*') {
				t.skipComment()
				continue
			}
		}

		// symbol
		if IsSymbol(buf[0]) {
			return &tokenInfo{
				TokenType: "SYMBOL",
				Symbol:    buf[0],
			}
		}

		// integer constant
		if IsNumber(buf[0]) {
			token += string(buf[0])
			for {
				la, ok := t.r.Lookahead()
				if !ok {
					break
				}

				if IsNumber(la) {
					utils.Must2(t.r.Read(buf))
					token += string(buf[0])
				} else {
					break
				}
			}

			// check whether integer is valid (0 <= x <= 32767)
			i, err := strconv.ParseInt(token, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("invalid number: %s, %+v", token, err))
			}
			if i < 0 || 32767 < i {
				panic(fmt.Sprintf("invalid number: %s, overflow", token))
			}

			return &tokenInfo{
				TokenType: "INT_CONST",
				IntVal:    int(i),
			}
		}

		// string constant
		if buf[0] == byte('"') {
			token := ""

			for {
				la, ok := t.r.Lookahead()
				if !ok {
					panic("string isn't finished")
				}

				if la == byte('\n') {
					panic("string cannot contain linebreak '\n'")
				}

				if la == byte('"') {
					utils.Must2(t.r.Read(buf))
					break
				}

				utils.Must2(t.r.Read(buf))
				token += string(buf[0])
			}

			return &tokenInfo{
				TokenType: "STRING_CONST",
				StringVal: token,
			}
		}

		// identifier or keyword
		if IsAlpha(buf[0]) || IsUnderscore(buf[0]) {
			// identifier
			token += string(buf[0])
			for {
				la, ok := t.r.Lookahead()
				if !ok {
					if IsKeyword(token) {
						return &tokenInfo{
							TokenType: "KEYWORD",
							Keyword:   token,
						}
					}
					return &tokenInfo{
						TokenType:  "IDENTIFIER",
						Identifier: token,
					}
				}

				if IsAlnum(la) || IsUnderscore(la) {
					utils.Must2(t.r.Read(buf))
					token += string(buf[0])
					continue
				} else {
					if IsKeyword(token) {
						return &tokenInfo{
							TokenType: "KEYWORD",
							Keyword:   token,
						}
					}
					return &tokenInfo{
						TokenType:  "IDENTIFIER",
						Identifier: token,
					}
				}
			}
		}
	}
}

func (t *Tokenizer) skipComment() {
	la, ok := t.r.Lookahead()
	if !ok {
		panic("unexpected error")
	}

	buf := make([]byte, 1)

	utils.Must2(t.r.Read(buf))

	// single line comment
	if la == '/' {

		for {
			_, err := t.r.Read(buf)
			if err != nil {
				if err == io.EOF {
					return
				}

				panic(fmt.Sprintf("io error: %+v", err))
			}

			if buf[0] == '\n' {
				return
			}
		}
	}

	// multi line comment
	if la == '*' {
		for {
			_, err := t.r.Read(buf)
			if err != nil {
				if err == io.EOF {
					return
				}

				panic(fmt.Sprintf("io error: %+v", err))
			}

			if buf[0] == '*' {
				b, ok := t.r.Lookahead()
				if !ok {
					panic("multi line comment is not closed, or io error occured")
				}

				if b == '/' {
					utils.Must2(t.r.Read(buf))
					return
				}
			}
		}
	}
}

func (t *Tokenizer) TokenType() string {
	return t.currentToken.TokenType
}

func (t *Tokenizer) KeyWord() string {
	return t.currentToken.Keyword
}

func (t *Tokenizer) Symbol() string {
	if t.currentToken.Symbol == '<' {
		return "&lt;"
	} else if t.currentToken.Symbol == '>' {
		return "&gt;"
	} else if t.currentToken.Symbol == '&' {
		return "&amp;"
	}
	return string(t.currentToken.Symbol)
}

func (t *Tokenizer) RawSymbol() byte {
	return t.currentToken.Symbol
}

func (t *Tokenizer) Identifier() string {
	return t.currentToken.Identifier
}

func (t *Tokenizer) IntVal() int {
	return t.currentToken.IntVal
}

func (t *Tokenizer) StringVal() string {
	return t.currentToken.StringVal
}
