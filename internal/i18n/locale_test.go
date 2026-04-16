package i18n

import (
	"testing"
)

func TestSetLanguage(t *testing.T) {
	tests := []struct {
		name string
		lang string
		want string
	}{
		{"set english", "en", "en"},
		{"set chinese", "zh", "zh"},
		{"set japanese", "ja", "ja"},
		{"invalid language", "invalid", "en"},
		{"empty language", "", "en"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLanguage(tt.lang)
			lang := GetCurrentLanguage()
			if lang != tt.want {
				t.Errorf("GetCurrentLanguage() = %v, want %v", lang, tt.want)
			}
		})
	}
}

func TestT(t *testing.T) {
	SetLanguage("en")

	tests := []struct {
		key  string
		want string
	}{
		{"root_use", "ohmypassword - A secure password generator"},
		{"generate_use", "Generate a new password"},
		{"version_use", "Show version information"},
		{"flag_length", "Password length (8-128, default: 16)"},
		{"unknown_key", "unknown_key"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			result := T(tt.key)
			if result != tt.want {
				t.Errorf("T(%q) = %v, want %v", tt.key, result, tt.want)
			}
		})
	}
}

func TestGetSupportedLanguages(t *testing.T) {
	languages := GetSupportedLanguages()

	if len(languages) == 0 {
		t.Error("GetSupportedLanguages() returned empty slice")
	}

	found := false
	for _, lang := range languages {
		if lang == "en" {
			found = true
			break
		}
	}
	if !found {
		t.Error("English language not found in supported languages")
	}
}

func TestTFormat(t *testing.T) {
	SetLanguage("en")
	result := TFormat("flag_length")
	if result == "" {
		t.Error("TFormat() should return non-empty string")
	}
}