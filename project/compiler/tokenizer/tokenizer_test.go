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
	assert.Equal("constructor", tokenizer.KeyWord())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("piyo", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("KEYWORD", tokenizer.TokenType())
	assert.Equal("function", tokenizer.KeyWord())

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
	assert.Equal("+", tokenizer.Symbol())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("b", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("SYMBOL", tokenizer.TokenType())
	assert.Equal("=", tokenizer.Symbol())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("c", tokenizer.Identifier())

	assert.False(tokenizer.HasMoreTokens())
}

func TestTokenizer_Symbol_Divider(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString(`
a / b = c
`)
	tokenizer := NewTokenizer(buf)

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("a", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("SYMBOL", tokenizer.TokenType())
	assert.Equal("/", tokenizer.Symbol())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("b", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("SYMBOL", tokenizer.TokenType())
	assert.Equal("=", tokenizer.Symbol())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("c", tokenizer.Identifier())

	assert.False(tokenizer.HasMoreTokens())
}

func TestTokenizer_IntegerConst(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString(`
1 + b = 12345
`)
	tokenizer := NewTokenizer(buf)

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("INT_CONST", tokenizer.TokenType())
	assert.Equal(1, tokenizer.IntVal())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("SYMBOL", tokenizer.TokenType())
	assert.Equal("+", tokenizer.Symbol())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("b", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("SYMBOL", tokenizer.TokenType())
	assert.Equal("=", tokenizer.Symbol())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("INT_CONST", tokenizer.TokenType())
	assert.Equal(12345, tokenizer.IntVal())

	assert.False(tokenizer.HasMoreTokens())
}

func TestTokenizer_StringConst(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString(`
a = "some string"
"zzz"
`)
	tokenizer := NewTokenizer(buf)

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("IDENTIFIER", tokenizer.TokenType())
	assert.Equal("a", tokenizer.Identifier())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("SYMBOL", tokenizer.TokenType())
	assert.Equal("=", tokenizer.Symbol())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("STRING_CONST", tokenizer.TokenType())
	assert.Equal("some string", tokenizer.StringVal())

	assert.True(tokenizer.HasMoreTokens())
	tokenizer.Advance()
	assert.Equal("STRING_CONST", tokenizer.TokenType())
	assert.Equal("zzz", tokenizer.StringVal())

	assert.False(tokenizer.HasMoreTokens())
}
