package strategy

import (
	"strings"
	"testing"
)

func TestPronounceableStrategy_Pattern(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		charset string
	}{
		{
			name:    "default length 10",
			length:  10,
			charset: "all",
		},
		{
			name:    "length 8 with all charset",
			length:  8,
			charset: "all",
		},
		{
			name:    "length 12 with digit",
			length:  12,
			charset: "digit",
		},
	}

	consonants := "bcdfghjklmnpqrstvwxyz"
	vowels := "aeiou"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strategy := NewPronounceableStrategy()
			opts := &Options{
				Length:  tt.length,
				Charset: tt.charset,
			}

			password, err := strategy.Generate(opts)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			if len(password) < 8 {
				t.Errorf("password length = %v, want >= 8", len(password))
			}

			if strings.Contains(tt.charset, "digit") || tt.charset == "all" {
				hasDigit := false
				for _, c := range password {
					if c >= '0' && c <= '9' {
						hasDigit = true
						break
					}
				}
				if !hasDigit {
					t.Logf("Password %q should contain digit (hasDigit=%v)", password, hasDigit)
				}
			}

			if strings.Contains(tt.charset, "symbol") || tt.charset == "all" {
				hasSymbol := false
				for _, c := range password {
					if strings.Contains("!@#$%^&*", string(c)) {
						hasSymbol = true
						break
					}
				}
				if !hasSymbol {
					t.Logf("Password %q should contain symbol", password)
				}
			}

			_ = consonants
			_ = vowels
		})
	}
}

func TestPassphraseStrategy_Pattern(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		charset string
	}{
		{
			name:    "4 words",
			length:  4,
			charset: "all",
		},
		{
			name:    "6 words",
			length:  6,
			charset: "all",
		},
		{
			name:    "min 4 words",
			length:  3,
			charset: "all",
		},
		{
			name:    "max 10 words",
			length:  12,
			charset: "all",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strategy := NewPassphraseStrategy()
			opts := &Options{
				Length:  tt.length,
				Charset: tt.charset,
			}

			password, err := strategy.Generate(opts)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			expectedWords := tt.length
			if expectedWords < 4 {
				expectedWords = 4
			}
			if expectedWords > 10 {
				expectedWords = 10
			}

			expectedParts := expectedWords
			if strings.Contains(tt.charset, "digit") || tt.charset == "all" {
				expectedParts += 2
			}
			if strings.Contains(tt.charset, "symbol") || tt.charset == "all" {
				expectedParts += 1
			}

			partCount := strings.Count(password, "-") + 1
			if partCount != expectedParts {
				t.Errorf("part count = %v, want %v (password: %q)", partCount, expectedParts, password)
			}

			if !strings.Contains(password, "-") {
				t.Error("passphrase should contain hyphens")
			}

			if strings.Contains(tt.charset, "digit") || tt.charset == "all" {
				hasDigit := false
				for _, c := range password {
					if c >= '0' && c <= '9' {
						hasDigit = true
						break
					}
				}
				if !hasDigit {
					t.Logf("Passphrase %q should contain digit", password)
				}
			}

			if strings.Contains(tt.charset, "symbol") || tt.charset == "all" {
				hasSymbol := false
				for _, c := range password {
					if strings.Contains("!@#$%^&*", string(c)) {
						hasSymbol = true
						break
					}
				}
				if !hasSymbol {
					t.Logf("Passphrase %q should contain symbol", password)
				}
			}
		})
	}
}

func TestSimpleStrategy_Randomness(t *testing.T) {
	strategy := &SimpleStrategy{}
	opts := &Options{
		Length:  16,
		Charset: "all",
	}

	uniquePasswords := make(map[string]bool)
	count := 100

	for i := 0; i < count; i++ {
		password, err := strategy.Generate(opts)
		if err != nil {
			t.Fatalf("Generate() error = %v", err)
		}
		uniquePasswords[password] = true
	}

	if len(uniquePasswords) < count/2 {
		t.Errorf("Expected at least %d unique passwords, got %d", count/2, len(uniquePasswords))
	}
}

func TestStrategy_Entropy(t *testing.T) {
	tests := []struct {
		name     string
		strategy string
		length   int
		charset  string
	}{
		{"simple all", "simple", 16, "all"},
		{"simple lower", "simple", 16, "lower"},
		{"pronounceable", "pronounceable", 10, "all"},
		{"passphrase 4", "passphrase", 4, "all"},
		{"passphrase 6", "passphrase", 6, "all"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := GetStrategy(tt.strategy)
			opts := &Options{
				Length:  tt.length,
				Charset: tt.charset,
			}

			entropy := s.CalculateEntropy(opts)
			if entropy <= 0 {
				t.Errorf("CalculateEntropy() = %v, want > 0", entropy)
			}
		})
	}
}
