package strategy

type Options struct {
	Length         int
	Charset        string
	Strategy       string
	Count          int
	Validate       bool
	Quiet          bool
	ExcludeSimilar bool
}

type Strategy interface {
	Generate(opts *Options) (string, error)
	CalculateEntropy(opts *Options) float64
}

var _ Strategy = (*SimpleStrategy)(nil)
var _ Strategy = (*PronounceableStrategy)(nil)
var _ Strategy = (*PassphraseStrategy)(nil)
var _ Strategy = (*MemorableStrategy)(nil)
var _ Strategy = (*SegmentedStrategy)(nil)

func GetStrategy(name string) Strategy {
	switch name {
	case "pronounceable":
		return NewPronounceableStrategy()
	case "passphrase":
		return NewPassphraseStrategy()
	case "memorable":
		return NewMemorableStrategy()
	case "segmented":
		return NewSegmentedStrategy()
	default:
		return &SimpleStrategy{}
	}
}
