package charset

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