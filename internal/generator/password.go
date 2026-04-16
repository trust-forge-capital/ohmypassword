package generator

import (
	"github.com/trust-forge-capital/ohmypassword/internal/random"
	"github.com/trust-forge-capital/ohmypassword/internal/strategy"
)

func GeneratePasswords(opts *Options) ([]string, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	strat := strategy.GetStrategy(opts.Strategy)
	charset := parseCharset(opts.Charset)

	opts.Charset = charset

	var passwords []string
	for i := 0; i < opts.Count; i++ {
		pwd, err := strat.Generate(opts)
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