package strategy

import (
	"github.com/trust-forge-capital/ohmypassword/internal/generator"
)

type Strategy interface {
	Generate(opts *generator.Options) (string, error)
	CalculateEntropy(opts *generator.Options) float64
}

func GetStrategy(name string) Strategy {
	switch name {
	case "pronounceable":
		return &PronounceableStrategy{}
	case "passphrase":
		return &PassphraseStrategy{}
	default:
		return &SimpleStrategy{}
	}
}