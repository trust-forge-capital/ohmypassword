package strategy

import (
	"github.com/trust-forge-capital/ohmypassword/internal/random"
	"github.com/trust-forge-capital/ohmypassword/pkg/charset"
)

type CustomStrategy struct {
	rules []Rule
	rng   random.RNG
}

type Rule struct {
	Type     string
	Charset  string
	MinCount int
	MaxCount int
	Position string
}

func NewCustomStrategy(rules []Rule) *CustomStrategy {
	return &CustomStrategy{
		rules: rules,
		rng:   random.NewCryptoRNG(),
	}
}

func (s *CustomStrategy) Generate(opts *Options) (string, error) {
	if len(s.rules) == 0 {
		simple := &SimpleStrategy{}
		return simple.Generate(opts)
	}

	result := make([]rune, 0, opts.Length)

	for _, rule := range s.rules {
		chars := charset.GetCharsetRunes(rule.Charset)
		count := rule.MinCount
		if rule.MaxCount > rule.MinCount {
			n, _ := s.rng.Intn(rule.MaxCount - rule.MinCount + 1)
			count = rule.MinCount + n
		}

		for i := 0; i < count && len(result) < opts.Length; i++ {
			idx, _ := s.rng.Intn(len(chars))
			result = append(result, chars[idx])
		}
	}

	for len(result) < opts.Length {
		chars := charset.GetCharsetRunes(opts.Charset)
		idx, _ := s.rng.Intn(len(chars))
		result = append(result, chars[idx])
	}

	return string(result), nil
}

func (s *CustomStrategy) CalculateEntropy(opts *Options) float64 {
	charsetSize := charset.GetCharsetSize(opts.Charset)
	return calculateEntropyBits(opts.Length, charsetSize)
}
