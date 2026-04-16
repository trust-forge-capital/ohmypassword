package strategy

import (
	"math"
	"strings"

	"github.com/trust-forge-capital/ohmypassword/internal/random"
)

type MemorableStrategy struct {
	consonants string
	vowels     string
	digits     string
	symbols    string
	rng        random.RNG
}

func NewMemorableStrategy() *MemorableStrategy {
	return &MemorableStrategy{
		consonants: "bcdfghjklmnpqrstvwxyz",
		vowels:     "aeiou",
		digits:     "0123456789",
		symbols:    "!@#$%^&*",
		rng:        random.NewCryptoRNG(),
	}
}

func (s *MemorableStrategy) Generate(opts *Options) (string, error) {
	length := opts.Length
	if length < 4 {
		length = 8
	}

	result := make([]rune, 0, length)

	hasDigit := strings.Contains(opts.Charset, "digit") || opts.Charset == "all"
	hasSymbol := strings.Contains(opts.Charset, "symbol") || opts.Charset == "all"

	for len(result) < length {
		c, err := s.randomConsonant()
		if err != nil {
			return "", err
		}
		result = append(result, c)

		if len(result) >= length {
			break
		}

		v, err := s.randomVowel()
		if err != nil {
			return "", err
		}
		result = append(result, v)

		if len(result) >= length {
			break
		}

		c2, err := s.randomConsonant()
		if err != nil {
			return "", err
		}
		result = append(result, c2)
	}

	if hasDigit {
		d, err := s.rng.Intn(len(s.digits))
		if err != nil {
			return "", err
		}
		result = append(result, rune(s.digits[d]))
	}

	if hasSymbol {
		sym, err := s.rng.Intn(len(s.symbols))
		if err != nil {
			return "", err
		}
		result = append(result, rune(s.symbols[sym]))
	}

	return string(result[:min(len(result), opts.Length)]), nil
}

func (s *MemorableStrategy) randomConsonant() (rune, error) {
	n, err := s.rng.Intn(len(s.consonants))
	if err != nil {
		return 0, err
	}
	return rune(s.consonants[n]), nil
}

func (s *MemorableStrategy) randomVowel() (rune, error) {
	n, err := s.rng.Intn(len(s.vowels))
	if err != nil {
		return 0, err
	}
	return rune(s.vowels[n]), nil
}

func (s *MemorableStrategy) CalculateEntropy(opts *Options) float64 {
	charsetSize := 21
	return float64(opts.Length) * math.Log2(float64(charsetSize))
}
