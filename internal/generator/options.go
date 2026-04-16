package generator

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
	if o.Strategy != "simple" && o.Strategy != "pronounceable" && o.Strategy != "passphrase" && o.Strategy != "memorable" {
		return ErrInvalidStrategy
	}
	if o.Charset != "all" && o.Charset != "upper" && o.Charset != "lower" &&
		o.Charset != "digit" && o.Charset != "symbol" &&
		o.Charset != "upper,lower" && o.Charset != "upper,lower,digit" &&
		o.Charset != "upper,lower,digit,symbol" {
		return ErrInvalidCharset
	}
	return nil
}
