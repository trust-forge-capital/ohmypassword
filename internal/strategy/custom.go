package strategy

import (
	"github.com/trust-forge-capital/ohmypassword/internal/generator"
	"github.com/trust-forge-capital/ohmypassword/internal/random"
)

type CustomStrategy struct {
	rules    []Rule
	rng      random.RNG
}

type Rule struct {
	Type      string
	Charset   string
	MinCount  int
	MaxCount  int
	Position  string
}

func NewCustomStrategy(rules []Rule) *CustomStrategy {
	return &CustomStrategy{
		rules: rules,
		rng:   random.NewCryptoRNG(),
	}
}

func (s *CustomStrategy) Generate(opts *generator.Options) (string, error) {
	if len(s.rules) == 0 {
		simple := &SimpleStrategy{}
		return simple.Generate(opts)
	}

	result := make([]rune, 0, opts.Length)

	for _, rule := range s.rules {
		charset := generator.GetCharset(rule.Charset)
		count := rule.MinCount
		if rule.MaxCount > rule.MinCount {
			n, _ := s.rng.Intn(rule.MaxCount - rule.MinCount + 1)
			count = rule.MinCount + n
		}

		for i := 0; i < count && len(result) < opts.Length; i++ {
			idx, _ := s.rng.Intn(len(charset))
			result = append(result, charset[idx])
		}
	}

	for len(result) < opts.Length {
		charset := generator.GetCharset(opts.Charset)
		idx, _ := s.rng.Intn(len(charset))
		result = append(result, charset[idx])
	}

	return string(result), nil
}

func (s *CustomStrategy) CalculateEntropy(opts *generator.Options) float64 {
	charsetSize := generator.GetCharsetSize(opts.Charset)
	return generator.CalculateEntropyBits(opts.Length, charsetSize)
}