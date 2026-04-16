package strategy

import (
	"github.com/trust-forge-capital/ohmypassword/internal/generator"
)

type SimpleStrategy struct{}

func (s *SimpleStrategy) Generate(opts *Options) (string, error) {
	charset := generator.GetCharset(opts.Charset)
	if opts.ExcludeSimilar {
		charset = generator.ExcludeSimilarChars(charset)
	}
	return generator.GenerateWithCharset(&generator.Options{
		Length:         opts.Length,
		Charset:        opts.Charset,
		Count:          opts.Count,
		Validate:       opts.Validate,
		Quiet:          opts.Quiet,
		ExcludeSimilar: opts.ExcludeSimilar,
	}, charset)
}

func (s *SimpleStrategy) CalculateEntropy(opts *Options) float64 {
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
