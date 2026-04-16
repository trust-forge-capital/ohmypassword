package generator

import "errors"

var (
	ErrInvalidLength   = errors.New("invalid length: must be between 8 and 128")
	ErrInvalidCount    = errors.New("invalid count: must be between 1 and 100")
	ErrInvalidStrategy = errors.New("invalid strategy: must be simple, pronounceable, or passphrase")
	ErrInvalidCharset  = errors.New("invalid charset")
	ErrNoCharacters    = errors.New("no characters available for generation")
)

type Password struct {
	Value  string
	Entropy float64
}

type Generator interface {
	Generate(opts *Options) (string, error)
	CalculateEntropy(opts *Options) float64
}