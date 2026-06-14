package charset

import "strings"

const (
	CharsetUpper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetLower  = "abcdefghijklmnopqrstuvwxyz"
	CharsetDigit  = "0123456789"
	CharsetSymbol = "!@#$%^&*()_+-=[]{}|;:,.<>?/~`\"'"
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
	return &UpperCharset{&BaseCharset{chars: CharsetUpper}}
}

type LowerCharset struct{ *BaseCharset }

func NewLowerCharset() *LowerCharset {
	return &LowerCharset{&BaseCharset{chars: CharsetLower}}
}

type DigitCharset struct{ *BaseCharset }

func NewDigitCharset() *DigitCharset {
	return &DigitCharset{&BaseCharset{chars: CharsetDigit}}
}

type SymbolCharset struct{ *BaseCharset }

func NewSymbolCharset() *SymbolCharset {
	return &SymbolCharset{&BaseCharset{chars: CharsetSymbol}}
}

type AllCharset struct{ *BaseCharset }

func NewAllCharset() *AllCharset {
	return &AllCharset{&BaseCharset{
		chars: CharsetUpper + CharsetLower + CharsetDigit + CharsetSymbol,
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

var defaultCharsetRunes = []rune(CharsetUpper + CharsetLower + CharsetDigit + CharsetSymbol)

var charsetRuneCache = map[string][]rune{
	"upper":                    []rune(CharsetUpper),
	"lower":                    []rune(CharsetLower),
	"digit":                    []rune(CharsetDigit),
	"lower,digit":              []rune(CharsetLower + CharsetDigit),
	"symbol":                   []rune(CharsetSymbol),
	"upper,lower":              []rune(CharsetUpper + CharsetLower),
	"upper,lower,digit":        []rune(CharsetUpper + CharsetLower + CharsetDigit),
	"upper,lower,digit,symbol": defaultCharsetRunes,
}

// GetCharsetRunes returns the rune set for the given charset name. The returned
// slice is a shared, pre-computed cache entry and must be treated as read-only;
// callers that need to mutate should copy it first.
func GetCharsetRunes(charset string) []rune {
	if chars, ok := charsetRuneCache[charset]; ok {
		return chars
	}
	return defaultCharsetRunes
}

func GetCharsetSize(charset string) int {
	return len(GetCharsetRunes(charset))
}

func IsValidCharset(charset string) bool {
	switch charset {
	case "upper", "lower", "digit", "lower,digit", "symbol",
		"upper,lower", "upper,lower,digit", "upper,lower,digit,symbol", "all":
		return true
	default:
		return false
	}
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

func GetExcludedSimilarCount(charset string) int {
	chars := GetCharsetRunes(charset)
	seen := make(map[rune]struct{})
	count := 0
	for _, c := range chars {
		if _, ok := SimilarChars[c]; ok {
			if _, already := seen[c]; !already {
				seen[c] = struct{}{}
				count++
			}
		}
	}
	return count
}

func DetectCharset(password string) string {
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSymbol := false

	for _, r := range password {
		switch {
		case r >= 'A' && r <= 'Z':
			hasUpper = true
		case r >= 'a' && r <= 'z':
			hasLower = true
		case r >= '0' && r <= '9':
			hasDigit = true
		default:
			if strings.ContainsRune(CharsetSymbol, r) {
				hasSymbol = true
			}
		}
	}

	var parts []string
	if hasUpper {
		parts = append(parts, "upper")
	}
	if hasLower {
		parts = append(parts, "lower")
	}
	if hasDigit {
		parts = append(parts, "digit")
	}
	if hasSymbol {
		parts = append(parts, "symbol")
	}

	if len(parts) == 0 {
		return "all"
	}
	return strings.Join(parts, ",")
}
