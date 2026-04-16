package validator

import (
	"testing"
)

func TestCalculateStrength(t *testing.T) {
	tests := []struct {
		name       string
		password   string
		charset    string
		wantLevel  string
		wantScore  int
	}{
		{
			name:       "strong password",
			password:   "aB3$kL9@mN2pQ!xY7",
			charset:    "all",
			wantLevel:  "strong",
			wantScore:  4,
		},
		{
			name:       "very strong password",
			password:   "xK9mP2nL5!qR8wT3yZ1@aB7",
			charset:    "all",
			wantLevel:  "very_strong",
			wantScore:  5,
		},
		{
			name:       "weak password",
			password:   "abc123",
			charset:    "lower,digit",
			wantLevel:  "weak",
			wantScore:  2,
		},
		{
			name:       "short password",
			password:   "abc",
			charset:    "lower",
			wantLevel:  "very_weak",
			wantScore:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateStrength(tt.password, tt.charset)
			if result.Level != tt.wantLevel {
				t.Errorf("Level = %v, want %v", result.Level, tt.wantLevel)
			}
			if result.Score != tt.wantScore {
				t.Errorf("Score = %v, want %v", result.Score, tt.wantScore)
			}
			if result.Entropy <= 0 {
				t.Error("Entropy should be > 0")
			}
		})
	}
}

func TestIsCommonPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"common password", "password", true},
		{"common 123456", "123456", true},
		{"qwerty", "qwerty", true},
		{"random", "random", false},
		{"random password", "Tr0ub4dor&3", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCommonPassword(tt.password)
			if result != tt.want {
				t.Errorf("IsCommonPassword() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestCalculateScore(t *testing.T) {
	tests := []struct {
		entropyBits int
		want        int
	}{
		{20, 1},
		{30, 2},
		{40, 3},
		{65, 4},
		{85, 5},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			score := calculateScore(tt.entropyBits)
			if score != tt.want {
				t.Errorf("calculateScore(%v) = %v, want %v", tt.entropyBits, score, tt.want)
			}
		})
	}
}