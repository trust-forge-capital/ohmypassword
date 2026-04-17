package generator

import (
	"math"
)

func CalculateEntropy(password string, charsetSize int) float64 {
	if charsetSize <= 0 || len(password) == 0 {
		return 0
	}
	return float64(len(password)) * math.Log2(float64(charsetSize))
}

func CalculateEntropyBits(length int, charsetSize int) int {
	if charsetSize <= 0 || length == 0 {
		return 0
	}
	entropy := float64(length) * math.Log2(float64(charsetSize))
	return int(math.Floor(entropy))
}

func IsEntropySufficient(password string, charsetSize int, minBits int) bool {
	entropy := CalculateEntropyBits(len(password), charsetSize)
	return entropy >= minBits
}

func GetEntropyLevel(entropyBits int) string {
	switch {
	case entropyBits < 28:
		return "very_weak"
	case entropyBits < 36:
		return "weak"
	case entropyBits < 60:
		return "reasonable"
	case entropyBits < 80:
		return "strong"
	default:
		return "very_strong"
	}
}

func EstimateCrackTime(entropyBits int) string {
	attackRates := []int{100, 1000, 1e6, 1e10, 1e12}

	combinations := math.Pow(2, float64(entropyBits))

	for _, rate := range attackRates {
		seconds := combinations / float64(rate)
		if seconds < 60 {
			return "< 1 second"
		} else if seconds < 3600 {
			return "< 1 hour"
		} else if seconds < 86400 {
			return "< 1 day"
		} else if seconds < 31536000 {
			return formatDuration(seconds)
		} else if seconds < 31536000*100 {
			return formatDuration(seconds)
		} else if seconds < 31536000*1000 {
			return "centuries"
		}
	}
	return "millennia+"
}

func formatDuration(seconds float64) string {
	years := int(seconds / 31536000)
	if years < 1 {
		days := int(seconds / 86400)
		return formatDays(days)
	}
	return formatYears(years)
}

func formatYears(years int) string {
	switch {
	case years < 10:
		return "< 10 years"
	case years < 100:
		return "< 100 years"
	case years < 1000:
		return "< 1,000 years"
	default:
		return "centuries"
	}
}

func formatDays(days int) string {
	switch {
	case days < 7:
		return "< 1 week"
	case days < 30:
		return "< 1 month"
	case days < 365:
		return "< 1 year"
	default:
		return formatYears(days / 365)
	}
}
