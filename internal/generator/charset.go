package generator

const (
	CharsetUpper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetLower  = "abcdefghijklmnopqrstuvwxyz"
	CharsetDigit  = "0123456789"
	CharsetSymbol = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

var SimilarChars = map[rune]rune{
	'0': 'O',
	'O': '0',
	'1': 'l',
	'l': '1',
	'I': '|',
	'|': 'I',
}

func GetCharset(charset string) []rune {
	var chars []rune

	switch charset {
	case "upper":
		chars = []rune(CharsetUpper)
	case "lower":
		chars = []rune(CharsetLower)
	case "digit":
		chars = []rune(CharsetDigit)
	case "lower,digit":
		chars = []rune(CharsetLower + CharsetDigit)
	case "symbol":
		chars = []rune(CharsetSymbol)
	case "upper,lower":
		chars = []rune(CharsetUpper + CharsetLower)
	case "upper,lower,digit":
		chars = []rune(CharsetUpper + CharsetLower + CharsetDigit)
	case "upper,lower,digit,symbol":
		chars = []rune(CharsetUpper + CharsetLower + CharsetDigit + CharsetSymbol)
	default:
		chars = []rune(CharsetUpper + CharsetLower + CharsetDigit + CharsetSymbol)
	}

	return chars
}

func GetCharsetSize(charset string) int {
	return len(GetCharset(charset))
}

func ExcludeSimilarChars(chars []rune) []rune {
	result := make([]rune, 0, len(chars))
	for _, c := range chars {
		if _, ok := SimilarChars[c]; !ok {
			result = append(result, c)
		}
	}
	return result
}

func GetExcludedSimilarCount() int {
	return len(SimilarChars) / 2
}
