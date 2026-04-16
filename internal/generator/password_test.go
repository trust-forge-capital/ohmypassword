package generator

import (
	"testing"
)

func TestGeneratePasswords(t *testing.T) {
	tests := []struct {
		name    string
		opts    *Options
		wantLen int
		wantErr bool
	}{
		{
			name:    "default options",
			opts:    &Options{Length: 16, Charset: "all", Strategy: "simple", Count: 1},
			wantLen: 16,
			wantErr: false,
		},
		{
			name:    "custom length",
			opts:    &Options{Length: 32, Charset: "all", Strategy: "simple", Count: 1},
			wantLen: 32,
			wantErr: false,
		},
		{
			name:    "lowercase only",
			opts:    &Options{Length: 12, Charset: "lower", Strategy: "simple", Count: 1},
			wantLen: 12,
			wantErr: false,
		},
		{
			name:    "digits only",
			opts:    &Options{Length: 8, Charset: "digit", Strategy: "simple", Count: 1},
			wantLen: 8,
			wantErr: false,
		},
		{
			name:    "invalid length too short",
			opts:    &Options{Length: 5, Charset: "all", Strategy: "simple", Count: 1},
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "invalid length too long",
			opts:    &Options{Length: 200, Charset: "all", Strategy: "simple", Count: 1},
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "invalid count",
			opts:    &Options{Length: 16, Charset: "all", Strategy: "simple", Count: 0},
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "multiple passwords",
			opts:    &Options{Length: 16, Charset: "all", Strategy: "simple", Count: 5},
			wantLen: 16,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passwords, err := GeneratePasswords(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePasswords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(passwords) != tt.opts.Count {
					t.Errorf("GeneratePasswords() count = %v, want %v", len(passwords), tt.opts.Count)
				}
				for _, pwd := range passwords {
					if len(pwd) != tt.wantLen {
						t.Errorf("GeneratePasswords() password length = %v, want %v", len(pwd), tt.wantLen)
					}
				}
			}
		})
	}
}

func TestOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    *Options
		wantErr bool
	}{
		{
			name:    "valid options",
			opts:    &Options{Length: 16, Charset: "all", Strategy: "simple", Count: 1},
			wantErr: false,
		},
		{
			name:    "invalid length",
			opts:    &Options{Length: 5, Charset: "all", Strategy: "simple", Count: 1},
			wantErr: true,
		},
		{
			name:    "invalid count",
			opts:    &Options{Length: 16, Charset: "all", Strategy: "simple", Count: 200},
			wantErr: true,
		},
		{
			name:    "invalid strategy",
			opts:    &Options{Length: 16, Charset: "all", Strategy: "invalid", Count: 1},
			wantErr: true,
		},
		{
			name:    "invalid charset",
			opts:    &Options{Length: 16, Charset: "invalid", Strategy: "simple", Count: 1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Options.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCharset(t *testing.T) {
	tests := []struct {
		name           string
		charset        string
		wantSize       int
		wantContains   string
	}{
		{"upper", "upper", 26, "A"},
		{"lower", "lower", 26, "a"},
		{"digit", "digit", 10, "5"},
		{"symbol", "symbol", 32, "!"},
		{"all", "all", 94, "A"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chars := GetCharset(tt.charset)
			if len(chars) != tt.wantSize {
				t.Errorf("GetCharset() size = %v, want %v", len(chars), tt.wantSize)
			}
			found := false
			for _, c := range chars {
				if string(c) == tt.wantContains {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("GetCharset() does not contain %v", tt.wantContains)
			}
		})
	}
}

func TestExcludeSimilarChars(t *testing.T) {
	tests := []struct {
		name       string
		charset    []rune
		wantSize   int
		shouldMiss string
	}{
		{
			name:       "exclude similar from all",
			charset:    []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*"),
			wantSize:   88,
			shouldMiss: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExcludeSimilarChars(tt.charset)
			if len(result) != tt.wantSize {
				t.Errorf("ExcludeSimilarChars() size = %v, want %v", len(result), tt.wantSize)
			}
			for _, c := range result {
				if string(c) == tt.shouldMiss {
					t.Errorf("ExcludeSimilarChars() should not contain %v", tt.shouldMiss)
				}
			}
		})
	}
}