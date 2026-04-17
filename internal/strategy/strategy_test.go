package strategy

import (
	"strings"
	"testing"
)

func TestSimpleStrategy_Generate(t *testing.T) {
	strategy := &SimpleStrategy{}
	opts := &Options{
		Length:  16,
		Charset: "all",
	}

	password, err := strategy.Generate(opts)
	if err != nil {
		t.Fatalf("SimpleStrategy.Generate() error = %v", err)
	}

	if len(password) != 16 {
		t.Errorf("password length = %v, want 16", len(password))
	}
}

func TestSimpleStrategy_CalculateEntropy(t *testing.T) {
	strategy := &SimpleStrategy{}
	opts := &Options{
		Length:  16,
		Charset: "all",
	}

	entropy := strategy.CalculateEntropy(opts)
	if entropy <= 0 {
		t.Errorf("CalculateEntropy() = %v, want > 0", entropy)
	}
}

func TestPronounceableStrategy_Generate(t *testing.T) {
	strategy := NewPronounceableStrategy()
	opts := &Options{
		Length:  10,
		Charset: "all",
	}

	password, err := strategy.Generate(opts)
	if err != nil {
		t.Fatalf("PronounceableStrategy.Generate() error = %v", err)
	}

	if len(password) < 8 {
		t.Errorf("password length = %v, want >= 8", len(password))
	}
}

func TestPassphraseStrategy_Generate(t *testing.T) {
	strategy := NewPassphraseStrategy()
	opts := &Options{
		Length:  4,
		Charset: "all",
	}

	password, err := strategy.Generate(opts)
	if err != nil {
		t.Fatalf("PassphraseStrategy.Generate() error = %v", err)
	}

	if len(password) < 10 {
		t.Errorf("password length = %v, want >= 10", len(password))
	}

	containsHyphen := false
	for _, c := range password {
		if c == '-' {
			containsHyphen = true
			break
		}
	}
	if !containsHyphen {
		t.Error("passphrase should contain hyphens")
	}
}

func TestMemorableStrategy_Generate(t *testing.T) {
	strategy := NewMemorableStrategy()
	opts := &Options{
		Length:  10,
		Charset: "all",
	}

	password, err := strategy.Generate(opts)
	if err != nil {
		t.Fatalf("MemorableStrategy.Generate() error = %v", err)
	}

	if len(password) < 8 {
		t.Errorf("password length = %v, want >= 8", len(password))
	}
}

func TestSegmentedStrategy_Generate(t *testing.T) {
	strategy := NewSegmentedStrategy()
	opts := &Options{
		Length:  12,
		Charset: "all",
	}

	password, err := strategy.Generate(opts)
	if err != nil {
		t.Fatalf("SegmentedStrategy.Generate() error = %v", err)
	}

	parts := strings.Split(password, "-")
	if len(parts) != 4 {
		t.Errorf("expected 4 segments, got %d", len(parts))
	}
	for _, part := range parts {
		if len(part) != 3 {
			t.Errorf("expected segment length 3, got %d", len(part))
		}
	}
}

func TestGetStrategy(t *testing.T) {
	tests := []struct {
		name     string
		strategy string
	}{
		{"simple", "simple"},
		{"pronounceable", "pronounceable"},
		{"passphrase", "passphrase"},
		{"memorable", "memorable"},
		{"segmented", "segmented"},
		{"default", "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := GetStrategy(tt.strategy)
			if s == nil {
				t.Error("GetStrategy() returned nil")
			}
		})
	}
}
