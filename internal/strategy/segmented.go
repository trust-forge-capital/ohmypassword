package strategy

import (
	"strings"

	"github.com/trust-forge-capital/ohmypassword/pkg/charset"
)

type SegmentedStrategy struct {
	separator     string
	segmentLength int
}

func NewSegmentedStrategy() *SegmentedStrategy {
	return &SegmentedStrategy{
		separator:     "-",
		segmentLength: 3,
	}
}

func (s *SegmentedStrategy) Generate(opts *Options) (string, error) {
	chars := charset.GetCharsetRunes(opts.Charset)
	if opts.ExcludeSimilar {
		chars = charset.ExcludeSimilarChars(chars)
	}
	chars = excludeSeparator(chars, s.separator)

	totalChars := opts.Length
	if totalChars < s.segmentLength {
		totalChars = s.segmentLength * 4 // default 12 chars = 4 segments
	}

	// round up to nearest multiple of segmentLength
	remainder := totalChars % s.segmentLength
	if remainder != 0 {
		totalChars += s.segmentLength - remainder
	}

	password, err := generateWithCharset(totalChars, chars)
	if err != nil {
		return "", err
	}

	return s.segment(string(password)), nil
}

// excludeSeparator returns a copy of chars with every rune of sep removed so
// the generated password can never contain the segment separator, which would
// otherwise make the segmented output ambiguous.
func excludeSeparator(chars []rune, sep string) []rune {
	result := make([]rune, 0, len(chars))
	for _, c := range chars {
		if !strings.ContainsRune(sep, c) {
			result = append(result, c)
		}
	}
	return result
}

func (s *SegmentedStrategy) segment(password string) string {
	var segments []string
	for i := 0; i < len(password); i += s.segmentLength {
		end := i + s.segmentLength
		if end > len(password) {
			end = len(password)
		}
		segments = append(segments, password[i:end])
	}
	return strings.Join(segments, s.separator)
}

func (s *SegmentedStrategy) CalculateEntropy(opts *Options) float64 {
	charsetSize := charset.GetCharsetSize(opts.Charset)
	if opts.ExcludeSimilar {
		excluded := charset.GetExcludedSimilarCount(opts.Charset)
		charsetSize -= excluded
		if charsetSize < 0 {
			charsetSize = 0
		}
	}

	totalChars := opts.Length
	if totalChars < s.segmentLength {
		totalChars = s.segmentLength * 4
	}
	remainder := totalChars % s.segmentLength
	if remainder != 0 {
		totalChars += s.segmentLength - remainder
	}

	return calculateEntropyBits(totalChars, charsetSize)
}
