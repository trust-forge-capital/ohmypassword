package strategy

import (
	"github.com/trust-forge-capital/ohmypassword/internal/generator"
)

type SimpleStrategy struct{}

func (s *SimpleStrategy) Generate(opts *generator.Options) (string, error) {
	charset := generator.GetCharset(opts.Charset)
	if opts.ExcludeSimilar {
		charset = generator.ExcludeSimilarChars(charset)
	}
	return generator.GenerateWithCharset(opts, charset)
}

func (s *SimpleStrategy) CalculateEntropy(opts *generator.Options) float64 {
	charsetSize := generator.GetCharsetSize(opts.Charset)
	if opts.ExcludeSimilar {
		excluded := generator.GetExcludedSimilarCount()
		charsetSize -= excluded
		if charsetSize < 0 {
			charsetSize = 0
		}
	}
	return generator.CalculateEntropyBits(opts.Length, charsetSize)
}