package generator

import "github.com/trust-forge-capital/ohmypassword/pkg/charset"

type Options struct {
	Length         int
	Charset        string
	Strategy       string
	Count          int
	ShowStrength   bool
	Quiet          bool
	ExcludeSimilar bool
}

func (o *Options) Validate() error {
	if o.Strategy == "passphrase" {
		if o.Length < 4 || o.Length > 10 {
			return ErrInvalidLength
		}
	} else {
		if o.Length < 8 || o.Length > 128 {
			return ErrInvalidLength
		}
	}
	if o.Count < 1 || o.Count > 100 {
		return ErrInvalidCount
	}
	if o.Strategy != "simple" && o.Strategy != "pronounceable" && o.Strategy != "passphrase" && o.Strategy != "memorable" && o.Strategy != "segmented" {
		return ErrInvalidStrategy
	}
	if !charset.IsValidCharset(o.Charset) {
		return ErrInvalidCharset
	}
	return nil
}
