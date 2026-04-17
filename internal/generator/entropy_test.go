package generator

import (
	"math"
	"testing"
)

func TestCalculateEntropy(t *testing.T) {
	tests := []struct {
		name        string
		password    string
		charsetSize int
		wantEntropy float64
		tolerance   float64
	}{
		{
			name:        "16 chars all charset (93)",
			password:    "abcdefghijklmnop",
			charsetSize: 93,
			wantEntropy: 104.6,
			tolerance:   1.0,
		},
		{
			name:        "8 chars lowercase (26)",
			password:    "abcdefgh",
			charsetSize: 26,
			wantEntropy: 37.6,
			tolerance:   1.0,
		},
		{
			name:        "8 chars digit (10)",
			password:    "01234567",
			charsetSize: 10,
			wantEntropy: 26.7,
			tolerance:   1.0,
		},
		{
			name:        "8 chars all (93)",
			password:    "abcdefgh",
			charsetSize: 93,
			wantEntropy: 52.3,
			tolerance:   1.0,
		},
		{
			name:        "32 chars all (93)",
			password:    "abcdefghijklmnopqrstuvwxyz012345",
			charsetSize: 93,
			wantEntropy: 209.3,
			tolerance:   1.0,
		},
		{
			name:        "empty password",
			password:    "",
			charsetSize: 26,
			wantEntropy: 0,
			tolerance:   0,
		},
		{
			name:        "zero charset",
			password:    "abc",
			charsetSize: 0,
			wantEntropy: 0,
			tolerance:   0,
		},
		{
			name:        "negative charset",
			password:    "abc",
			charsetSize: -1,
			wantEntropy: 0,
			tolerance:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entropy := CalculateEntropy(tt.password, tt.charsetSize)
			if tt.tolerance == 0 {
				if entropy != tt.wantEntropy {
					t.Errorf("CalculateEntropy() = %v, want %v", entropy, tt.wantEntropy)
				}
			} else {
				diff := math.Abs(entropy - tt.wantEntropy)
				if diff > tt.tolerance {
					t.Errorf("CalculateEntropy() = %v, want ~%v (diff %v > tolerance %v)",
						entropy, tt.wantEntropy, diff, tt.tolerance)
				}
			}
		})
	}
}

func TestCalculateEntropyBits(t *testing.T) {
	tests := []struct {
		name        string
		length      int
		charsetSize int
		wantBits    int
	}{
		{
			name:        "16 chars all",
			length:      16,
			charsetSize: 93,
			wantBits:    104,
		},
		{
			name:        "8 chars lower",
			length:      8,
			charsetSize: 26,
			wantBits:    37,
		},
		{
			name:        "8 chars digit",
			length:      8,
			charsetSize: 10,
			wantBits:    26,
		},
		{
			name:        "zero length",
			length:      0,
			charsetSize: 26,
			wantBits:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bits := CalculateEntropyBits(tt.length, tt.charsetSize)
			if bits != tt.wantBits {
				t.Errorf("CalculateEntropyBits() = %v, want %v", bits, tt.wantBits)
			}
		})
	}
}

func TestGetEntropyLevel(t *testing.T) {
	tests := []struct {
		name      string
		entropy   int
		wantLevel string
	}{
		{"very weak - 27 bits", 27, "very_weak"},
		{"very weak - 10 bits", 10, "very_weak"},
		{"weak - 28 bits", 28, "weak"},
		{"weak - 35 bits", 35, "weak"},
		{"reasonable - 36 bits", 36, "reasonable"},
		{"reasonable - 59 bits", 59, "reasonable"},
		{"strong - 60 bits", 60, "strong"},
		{"strong - 79 bits", 79, "strong"},
		{"very strong - 80 bits", 80, "very_strong"},
		{"very strong - 100 bits", 100, "very_strong"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			level := GetEntropyLevel(tt.entropy)
			if level != tt.wantLevel {
				t.Errorf("GetEntropyLevel(%d) = %v, want %v", tt.entropy, level, tt.wantLevel)
			}
		})
	}
}

func TestEstimateCrackTime(t *testing.T) {
	tests := []struct {
		name         string
		entropyBits  int
		wantContains string
	}{
		{"very weak - 20 bits", 20, "day"},
		{"weak - 30 bits", 30, "year"},
		{"reasonable - 45 bits", 45, "years"},
		{"reasonable - 55 bits", 55, "year"},
		{"strong - 65 bits", 65, "centuries"},
		{"strong - 75 bits", 75, "millennia"},
		{"very strong - 85 bits", 85, "millennia"},
		{"very strong - 100 bits", 100, "millennia"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crackTime := EstimateCrackTime(tt.entropyBits)
			t.Logf("EstimateCrackTime(%d) = %v", tt.entropyBits, crackTime)
			if !contains(crackTime, tt.wantContains) {
				t.Errorf("EstimateCrackTime(%d) = %v, want to contain %v",
					tt.entropyBits, crackTime, tt.wantContains)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
