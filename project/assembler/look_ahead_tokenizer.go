package main

type LookAheadTokenizer struct {
	t         *Tokenizer
	lookAhead *Token
	err       error
}

func NewLookAheadTokenizer(t *Tokenizer) *LookAheadTokenizer {
	return &LookAheadTokenizer{t: t}
}

func (t *LookAheadTokenizer) Read() (*Token, error) {
	if t.lookAhead != nil {
		ret := t.lookAhead
		t.lookAhead = nil
		return ret, nil
	}

	token, err := t.t.Read()
	return &token, err
}

func (t *LookAheadTokenizer) LookAhead() (*Token, error) {
	if t.err != nil {
		return nil, t.err
	}

	if t.lookAhead != nil {
		return t.lookAhead, nil
	}

	token, err := t.t.Read()

	if err != nil {
		t.lookAhead = nil
		return nil, err
	}

	t.lookAhead = &token
	return &token, nil
}
