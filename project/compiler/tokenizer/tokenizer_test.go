package tokenizer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer_SkipComment(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString(`
	// skip whitespaces
// aaa
/* multi 

  line comment


*/


`)
	tokenizer := NewTokenizer(buf)
	assert.False(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.False(tokenizer.HasMoreTokens())
}

func TestTokenizer_Identifier(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString(`
hoge
	fuga piyo
`)
	tokenizer := NewTokenizer(buf)

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("hoge", tokenizer.Identifier())
	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("fuga", tokenizer.Identifier())
	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("piyo", tokenizer.Identifier())
	assert.False(tokenizer.HasMoreTokens())
}

func TestTokenizer_Keyword(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString(`
hoge
	constructor piyo function
`)
	tokenizer := NewTokenizer(buf)

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("hoge", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("KEYWORD", tokenizer.TokenType())
	assert.Equal("constructor", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("piyo", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("KEYWORD", tokenizer.TokenType())
	assert.Equal("function", tokenizer.Identifier())

	assert.False(tokenizer.HasMoreTokens())
}

func TestTokenizer_Symbol(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString(`
a + b = c
`)
	tokenizer := NewTokenizer(buf)

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("a", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("SYMBOL", tokenizer.TokenType())
	assert.Equal(byte('+'), tokenizer.Symbol())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("b", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("SYMBOL", tokenizer.TokenType())
	assert.Equal(byte('='), tokenizer.Symbol())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("c", tokenizer.Identifier())

	assert.False(tokenizer.HasMoreTokens())
}
