package tokenizer

import "io"

type LookaheadReader struct {
	r            io.Reader
	lookahead    []byte
	lookaheadErr error
	lookaheaded  bool
}

func NewLookaheadReader(r io.Reader) *LookaheadReader {
	return &LookaheadReader{
		r:           r,
		lookahead:   make([]byte, 1),
		lookaheaded: false,
	}
}

func (r *LookaheadReader) Read(buf []byte) (int, error) {
	if r.lookaheaded {
		if r.lookaheadErr != nil {
			return 0, r.lookaheadErr
		}

		r.lookaheaded = false
		buf[0] = r.lookahead[0]
		return 1, nil
	}

	return r.r.Read(buf)
}

func (r *LookaheadReader) Lookahead() (byte, bool) {
	if r.lookaheaded {
		return r.lookahead[0], r.lookaheadErr == nil
	}

	_, err := r.r.Read(r.lookahead)
	r.lookaheadErr = err
	r.lookaheaded = true
	return r.lookahead[0], err == nil
}
