package main

func IsAlpha(c rune) bool {
	return 'A' <= c && c <= 'z'
}

func IsNumber(c rune) bool {
	return '0' <= c && c <= '9'
}

func IsSymbol(c rune) bool {
	switch c {
	case '_':
		return true
	case '.':
		return true
	case '$':
		return true
	case ':':
		return true
	default:
		return false
	}

	return false
}
