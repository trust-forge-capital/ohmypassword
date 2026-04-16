package strategy

import (
	"math"
	"strings"

	"github.com/trust-forge-capital/ohmypassword/internal/generator"
	"github.com/trust-forge-capital/ohmypassword/internal/random"
)

type PronounceableStrategy struct {
	consonants string
	vowels     string
	digits     string
	rng        random.RNG
}

func NewPronounceableStrategy() *PronounceableStrategy {
	return &PronounceableStrategy{
		consonants: "bcdfghjklmnpqrstvwxyz",
		vowels:     "aeiou",
		digits:     "0123456789",
		rng:        random.NewCryptoRNG(),
	}
}

func (s *PronounceableStrategy) Generate(opts *Options) (string, error) {
	result := make([]rune, 0, opts.Length)
	length := opts.Length

	if length < 4 {
		length = 8
	}

	hasDigit := strings.Contains(opts.Charset, "digit") || opts.Charset == "all"
	hasSymbol := strings.Contains(opts.Charset, "symbol") || opts.Charset == "all"

	for len(result) < length-2 {
		c, err := s.randomConsonant()
		if err != nil {
			return "", err
		}
		result = append(result, c)

		if len(result) < length {
			v, err := s.randomVowel()
			if err != nil {
				return "", err
			}
			result = append(result, v)
		}
	}

	if hasDigit {
		d, _ := s.rng.Intn(len(s.digits))
		result = append(result, rune(s.digits[d]))
	}

	if hasSymbol {
		symbols := "!@#$%^&*"
		pos, _ := s.rng.Intn(len(result))
		sym, _ := s.rng.Intn(len(symbols))
		symbolsRune := []rune(symbols)
		result = insertRune(result, pos, symbolsRune[sym])
	}

	return string(result[:min(len(result), opts.Length)]), nil
}

func (s *PronounceableStrategy) randomConsonant() (rune, error) {
	n, err := s.rng.Intn(len(s.consonants))
	if err != nil {
		return 0, err
	}
	return rune(s.consonants[n]), nil
}

func (s *PronounceableStrategy) randomVowel() (rune, error) {
	n, err := s.rng.Intn(len(s.vowels))
	if err != nil {
		return 0, err
	}
	return rune(s.vowels[n]), nil
}

func (s *PronounceableStrategy) CalculateEntropy(opts *Options) float64 {
	charsetSize := 20
	return float64(opts.Length) * math.Log2(float64(charsetSize))
}

func insertRune(slice []rune, index int, value rune) []rune {
	result := make([]rune, len(slice)+1)
	copy(result, slice[:index])
	result[index] = value
	copy(result[index+1:], slice[index:])
	return result
}
