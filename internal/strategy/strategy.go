package strategy

import (
	"github.com/trust-forge-capital/ohmypassword/internal/generator"
)

type Strategy interface {
	Generate(opts *generator.Options) (string, error)
	CalculateEntropy(opts *generator.Options) float64
}

var _ Strategy = (*SimpleStrategy)(nil)
var _ Strategy = (*PronounceableStrategy)(nil)
var _ Strategy = (*PassphraseStrategy)(nil)

func GetStrategy(name string) Strategy {
	switch name {
	case "pronounceable":
		return NewPronounceableStrategy()
	case "passphrase":
		return NewPassphraseStrategy()
	default:
		return &SimpleStrategy{}
	}
}