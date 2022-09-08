package tokenizer

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookaheadReader(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBuffer([]byte("hoge"))
	r := NewLookaheadReader(buf)

	la, ok := r.Lookahead()
	assert.Equal(byte('h'), la)
	assert.True(ok)

	la, ok = r.Lookahead()
	assert.Equal(byte('h'), la)
	assert.True(ok)

	b := make([]byte, 1)
	_, err := r.Read(b)
	assert.Equal(byte('h'), b[0])
	assert.Nil(err)

	_, err = r.Read(b)
	assert.Equal(byte('o'), b[0])
	assert.Nil(err)

	la, ok = r.Lookahead()
	assert.Equal(byte('g'), la)
	assert.True(ok)

	_, err = r.Read(b)
	assert.Equal(byte('g'), b[0])
	assert.Nil(err)

	_, err = r.Read(b)
	assert.Equal(byte('e'), b[0])
	assert.Nil(err)

	_, err = r.Read(b)
	assert.Equal(io.EOF, err)

	_, ok = r.Lookahead()
	assert.False(ok)
}
