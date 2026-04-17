package validator

import (
	"regexp"
	"strings"

	"github.com/trust-forge-capital/ohmypassword/internal/generator"
	"github.com/trust-forge-capital/ohmypassword/pkg/charset"
)

type Validator interface {
	Validate(password string) (bool, string)
}

type StrengthResult struct {
	Level       string
	Entropy     float64
	CrackTime   string
	Score       int
	Suggestions []string
}

func CalculateStrength(password string, charsetName string) StrengthResult {
	charsetSize := charset.GetCharsetSize(charsetName)
	entropy := generator.CalculateEntropy(password, charsetSize)
	entropyBits := int(entropy)

	level := generator.GetEntropyLevel(entropyBits)
	crackTime := generator.EstimateCrackTime(entropyBits)
	score := calculateScore(entropyBits)
	suggestions := generateSuggestions(password, charsetSize)

	return StrengthResult{
		Level:       level,
		Entropy:     entropy,
		CrackTime:   crackTime,
		Score:       score,
		Suggestions: suggestions,
	}
}

func calculateScore(entropyBits int) int {
	switch {
	case entropyBits < 28:
		return 1
	case entropyBits < 36:
		return 2
	case entropyBits < 60:
		return 3
	case entropyBits < 80:
		return 4
	default:
		return 5
	}
}

func generateSuggestions(password string, charsetSize int) []string {
	var suggestions []string

	if len(password) < 12 {
		suggestions = append(suggestions, "Use a longer password (12+ characters)")
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSymbol := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{}|;:,.<>?]`).MatchString(password)

	if !hasLower {
		suggestions = append(suggestions, "Add lowercase letters")
	}
	if !hasUpper {
		suggestions = append(suggestions, "Add uppercase letters")
	}
	if !hasDigit {
		suggestions = append(suggestions, "Add numbers")
	}
	if !hasSymbol {
		suggestions = append(suggestions, "Add special characters")
	}

	commonPasswords := []string{"password", "123456", "qwerty", "admin", "letmein"}
	lower := strings.ToLower(password)
	for _, common := range commonPasswords {
		if strings.Contains(lower, common) {
			suggestions = append(suggestions, "Avoid common passwords")
			break
		}
	}

	return suggestions
}

func IsCommonPassword(password string) bool {
	commonPasswords := map[string]bool{
		"password": true, "123456": true, "12345678": true, "qwerty": true,
		"admin": true, "letmein": true, "welcome": true, "monkey": true,
		"dragon": true, "master": true, "1234567890": true, "abc123": true,
	}
	return commonPasswords[strings.ToLower(password)]
}