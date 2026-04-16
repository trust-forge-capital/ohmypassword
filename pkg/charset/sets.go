package charset

type CharsetOption func(*CompositeCharset)

type CompositeCharset struct {
	charsets []Charset
}

func NewCompositeCharset(options ...CharsetOption) *CompositeCharset {
	cc := &CompositeCharset{}
	for _, opt := range options {
		opt(cc)
	}
	return cc
}

func WithUpper() CharsetOption {
	return func(c *CompositeCharset) {
		c.charsets = append(c.charsets, NewUpperCharset())
	}
}

func WithLower() CharsetOption {
	return func(c *CompositeCharset) {
		c.charsets = append(c.charsets, NewLowerCharset())
	}
}

func WithDigits() CharsetOption {
	return func(c *CompositeCharset) {
		c.charsets = append(c.charsets, NewDigitCharset())
	}
}

func WithSymbols() CharsetOption {
	return func(c *CompositeCharset) {
		c.charsets = append(c.charsets, NewSymbolCharset())
	}
}

func (c *CompositeCharset) String() string {
	result := ""
	for _, cs := range c.charsets {
		result += cs.String()
	}
	return result
}

func (c *CompositeCharset) Size() int {
	result := 0
	for _, cs := range c.charsets {
		result += cs.Size()
	}
	return result
}

func (c *CompositeCharset) Contains(r rune) bool {
	for _, cs := range c.charsets {
		if cs.Contains(r) {
			return true
		}
	}
	return false
}

var (
	Upper   = NewUpperCharset()
	Lower   = NewLowerCharset()
	Digits  = NewDigitCharset()
	Symbols = NewSymbolCharset()
	All     = NewAllCharset()
)