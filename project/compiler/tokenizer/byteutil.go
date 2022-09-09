package tokenizer

var symbols = []byte{
	'{',
	'}',
	'(',
	')',
	'[',
	']',
	'.',
	',',
	';',
	'+',
	'-',
	'*',
	'/',
	'&',
	'|',
	'<',
	'>',
	'=',
	'~',
}

func IsSymbol(b byte) bool {
	for _, symbol := range symbols {
		if b == symbol {
			return true
		}
	}
	return false
}

func IsAlpha(b byte) bool {
	lowerAlphabet := 'a' <= b && b <= 'z'
	upperAlphabet := 'A' <= b && b <= 'Z'
	return lowerAlphabet || upperAlphabet
}

func IsNumber(b byte) bool {
	return '0' <= b && b <= '9'
}

func IsAlnum(b byte) bool {
	return IsAlpha(b) || IsNumber(b)
}

func IsUnderscore(b byte) bool {
	return b == '_'
}
