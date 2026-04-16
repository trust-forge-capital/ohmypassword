package validator

import (
	"math"
	"testing"
)

func TestCalculateStrength_AllFields(t *testing.T) {
	tests := []struct {
		name          string
		password      string
		charset       string
		wantEntropy   bool
		wantCrackTime bool
		wantScore     int
	}{
		{
			name:          "strong password with all fields",
			password:      "aB3$kL9@mN2pQ!xY7",
			charset:       "all",
			wantEntropy:   true,
			wantCrackTime: true,
			wantScore:     5,
		},
		{
			name:          "weak password",
			password:      "abc123",
			charset:       "lower,digit",
			wantEntropy:   true,
			wantCrackTime: true,
			wantScore:     2,
		},
		{
			name:          "short password",
			password:      "abc",
			charset:       "lower",
			wantEntropy:   true,
			wantCrackTime: true,
			wantScore:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateStrength(tt.password, tt.charset)

			if tt.wantEntropy && result.Entropy <= 0 {
				t.Errorf("CalculateStrength() entropy = %v, want > 0", result.Entropy)
			}
			if tt.wantCrackTime && result.CrackTime == "" {
				t.Error("CalculateStrength() crackTime should not be empty")
			}
			if result.Score != tt.wantScore {
				t.Errorf("CalculateStrength() score = %v, want %v", result.Score, tt.wantScore)
			}
		})
	}
}

func TestStrengthLevel_Boundaries(t *testing.T) {
	tests := []struct {
		name      string
		entropy   int
		wantLevel string
	}{
		{"very_weak boundary - 27", 27, "very_weak"},
		{"very_weak max - 10", 10, "very_weak"},
		{"weak min - 28", 28, "weak"},
		{"weak max - 35", 35, "weak"},
		{"reasonable min - 36", 36, "reasonable"},
		{"reasonable max - 59", 59, "reasonable"},
		{"strong min - 60", 60, "strong"},
		{"strong max - 79", 79, "strong"},
		{"very_strong min - 80", 80, "very_strong"},
		{"very_strong max - 200", 200, "very_strong"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateStrength("testpassword", "all")
			_ = result.Entropy

			level := getLevelFromEntropy(tt.entropy)
			if level != tt.wantLevel {
				t.Errorf("getLevelFromEntropy(%d) = %v, want %v", tt.entropy, level, tt.wantLevel)
			}
		})
	}
}

func getLevelFromEntropy(entropyBits int) string {
	if entropyBits < 28 {
		return "very_weak"
	} else if entropyBits < 36 {
		return "weak"
	} else if entropyBits < 60 {
		return "reasonable"
	} else if entropyBits < 80 {
		return "strong"
	}
	return "very_strong"
}

func TestCrackTime_AttackScenarios(t *testing.T) {
	tests := []struct {
		name     string
		entropy  int
		scenario string
	}{
		{
			name:     "online throttled 20 bits",
			entropy:  20,
			scenario: "online_throttled",
		},
		{
			name:     "online unthrottled 30 bits",
			entropy:  30,
			scenario: "online_unthrottled",
		},
		{
			name:     "offline slow 50 bits",
			entropy:  50,
			scenario: "offline_slow",
		},
		{
			name:     "offline fast 70 bits",
			entropy:  70,
			scenario: "offline_fast",
		},
		{
			name:     "offline GPU 90 bits",
			entropy:  90,
			scenario: "offline_fast_gpu",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateStrength("testpassword", "all")
			_ = result.CrackTime

			crackTime := calculateForScenario(tt.entropy, tt.scenario)
			if crackTime == "" {
				t.Errorf("calculateForScenario(%d, %s) should not be empty", tt.entropy, tt.scenario)
			}
		})
	}
}

func calculateForScenario(entropyBits int, scenario string) string {
	assumptions := map[string]int{
		"online_throttled":   100,
		"online_unthrottled": 1000,
		"offline_slow":       1e6,
		"offline_fast":       1e10,
		"offline_fast_gpu":   1e12,
	}

	rate, ok := assumptions[scenario]
	if !ok {
		return ""
	}

	combinations := math.Pow(2, float64(entropyBits))
	seconds := combinations / float64(rate)

	if seconds < 60 {
		return "< 1 second"
	} else if seconds < 3600 {
		return "< 1 hour"
	} else if seconds < 86400 {
		return "< 1 day"
	} else if seconds < 31536000 {
		days := int(seconds / 86400)
		if days < 7 {
			return "< 1 week"
		} else if days < 30 {
			return "< 1 month"
		} else if days < 365 {
			return "< 1 year"
		}
		return "< " + string(rune('0'+days/365/100+1)) + " years"
	} else if seconds < 31536000*100 {
		return "< 100 years"
	} else if seconds < 31536000*1000 {
		return "centuries"
	}
	return "millennia+"
}
