package charset

const (
	CharsetUpper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetLower  = "abcdefghijklmnopqrstuvwxyz"
	CharsetDigit  = "0123456789"
	CharsetSymbol = "!@#$%^&*()_+-=[]{}|;:,.<>?/~`-\""
)

var SimilarChars = map[rune]rune{
	'0': 'O',
	'O': '0',
	'1': 'l',
	'l': '1',
	'I': '|',
	'|': 'I',
}

type Charset interface {
	String() string
	Size() int
	Contains(r rune) bool
}

type BaseCharset struct {
	chars string
}

func (c *BaseCharset) String() string {
	return c.chars
}

func (c *BaseCharset) Size() int {
	return len(c.chars)
}

func (c *BaseCharset) Contains(r rune) bool {
	for _, c := range c.chars {
		if c == r {
			return true
		}
	}
	return false
}

type UpperCharset struct{ *BaseCharset }

func NewUpperCharset() *UpperCharset {
	return &UpperCharset{&BaseCharset{chars: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}}
}

type LowerCharset struct{ *BaseCharset }

func NewLowerCharset() *LowerCharset {
	return &LowerCharset{&BaseCharset{chars: "abcdefghijklmnopqrstuvwxyz"}}
}

type DigitCharset struct{ *BaseCharset }

func NewDigitCharset() *DigitCharset {
	return &DigitCharset{&BaseCharset{chars: "0123456789"}}
}

type SymbolCharset struct{ *BaseCharset }

func NewSymbolCharset() *SymbolCharset {
	return &SymbolCharset{&BaseCharset{chars: "!@#$%^&*()_+-=[]{}|;:,.<>?"}}
}

type AllCharset struct{ *BaseCharset }

func NewAllCharset() *AllCharset {
	return &AllCharset{&BaseCharset{
		chars: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|;:,.<>?",
	}}
}

func GetCharset(name string) Charset {
	switch name {
	case "upper":
		return NewUpperCharset()
	case "lower":
		return NewLowerCharset()
	case "digit":
		return NewDigitCharset()
	case "symbol":
		return NewSymbolCharset()
	default:
		return NewAllCharset()
	}
}

func GetCharsetRunes(charset string) []rune {
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
	return len(GetCharsetRunes(charset))
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
