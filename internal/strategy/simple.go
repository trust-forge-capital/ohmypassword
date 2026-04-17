package strategy

import (
	"errors"
	"math"

	"github.com/trust-forge-capital/ohmypassword/internal/random"
	"github.com/trust-forge-capital/ohmypassword/pkg/charset"
)

var ErrNoCharacters = errors.New("no characters available")

type SimpleStrategy struct{}

func (s *SimpleStrategy) Generate(opts *Options) (string, error) {
	chars := charset.GetCharsetRunes(opts.Charset)
	if opts.ExcludeSimilar {
		chars = charset.ExcludeSimilarChars(chars)
	}
	return generateWithCharset(opts.Length, chars)
}

func (s *SimpleStrategy) CalculateEntropy(opts *Options) float64 {
	charsetSize := charset.GetCharsetSize(opts.Charset)
	if opts.ExcludeSimilar {
		excluded := charset.GetExcludedSimilarCount(opts.Charset)
		charsetSize -= excluded
		if charsetSize < 0 {
			charsetSize = 0
		}
	}
	return calculateEntropyBits(opts.Length, charsetSize)
}

func generateWithCharset(length int, charSet []rune) (string, error) {
	if len(charSet) == 0 {
		return "", ErrNoCharacters
	}

	rng := random.NewCryptoRNG()
	result := make([]rune, length)

	for i := 0; i < length; i++ {
		n, err := rng.Intn(len(charSet))
		if err != nil {
			return "", err
		}
		result[i] = charSet[n]
	}

	return string(result), nil
}

func calculateEntropyBits(length int, charsetSize int) float64 {
	if charsetSize <= 0 || length == 0 {
		return 0
	}
	entropy := float64(length) * math.Log2(float64(charsetSize))
	return entropy
}
