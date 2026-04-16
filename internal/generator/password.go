package generator

import (
	"github.com/trust-forge-capital/ohmypassword/internal/random"
	"github.com/trust-forge-capital/ohmypassword/internal/strategy"
)

func GeneratePasswords(opts *Options) ([]string, error) {
	if err := opts.Validate(); err != nil {
		return nil, nil
	}

	strat := strategy.GetStrategy(opts.Strategy)
	charset := parseCharset(opts.Charset)

	opts.Charset = charset

	strategyOpts := &strategy.Options{
		Length:         opts.Length,
		Charset:        opts.Charset,
		Strategy:       opts.Strategy,
		Count:          opts.Count,
		Validate:       opts.Validate,
		Quiet:          opts.Quiet,
		ExcludeSimilar: opts.ExcludeSimilar,
	}

	var passwords []string
	for i := 0; i < opts.Count; i++ {
		pwd, err := strat.Generate(strategyOpts)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, pwd)
	}

	return passwords, nil
}

func parseCharset(charset string) string {
	if charset == "all" {
		return "upper,lower,digit,symbol"
	}
	return charset
}

func GenerateWithCharset(opts *Options, charSet []rune) (string, error) {
	if len(charSet) == 0 {
		return "", ErrNoCharacters
	}

	rng := random.NewCryptoRNG()
	result := make([]rune, opts.Length)

	for i := 0; i < opts.Length; i++ {
		n, err := rng.Intn(len(charSet))
		if err != nil {
			return "", err
		}
		result[i] = charSet[n]
	}

	return string(result), nil
}
