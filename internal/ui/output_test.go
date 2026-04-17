package ui

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/jedib0t/go-pretty/v6/table"
)

//nolint:errcheck
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	_ = w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestOutput_JSON(t *testing.T) {
	results := []PasswordResult{
		{Password: "abc123", Entropy: 50.5, Strength: StrengthInfo{Level: "Strong", CrackTime: "centuries", Score: 4}},
	}
	out := captureOutput(func() {
		if err := Output(results, "json", false); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	var parsed []map[string]any
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("failed to parse JSON: %v", err)
	}
	if len(parsed) != 1 {
		t.Fatalf("expected 1 result, got %d", len(parsed))
	}
	if parsed[0]["password"] != "abc123" {
		t.Errorf("expected password abc123, got %v", parsed[0]["password"])
	}
}

func TestOutput_CSV(t *testing.T) {
	results := []PasswordResult{
		{Password: "abc123", Entropy: 50.5, Strength: StrengthInfo{Level: "Strong", CrackTime: "centuries"}},
	}
	out := captureOutput(func() {
		if err := Output(results, "csv", false); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected header + 1 data row, got %d lines", len(lines))
	}
	if !strings.Contains(lines[0], "password") {
		t.Errorf("expected password header, got %s", lines[0])
	}
}

func TestOutput_Table(t *testing.T) {
	results := []PasswordResult{
		{Password: "abc123", Entropy: 50.5, Strength: StrengthInfo{Level: "Strong", CrackTime: "centuries"}},
	}
	out := captureOutput(func() {
		if err := Output(results, "table", false); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	if !strings.Contains(out, "abc123") {
		t.Errorf("expected table to contain password")
	}
}

func TestOutput_Simple(t *testing.T) {
	results := []PasswordResult{
		{Password: "abc123"},
	}
	out := captureOutput(func() {
		if err := Output(results, "simple", false); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	if !strings.Contains(out, "abc123") {
		t.Errorf("expected simple output to contain password")
	}
}

func TestOutput_Quiet(t *testing.T) {
	results := []PasswordResult{
		{Password: "abc123", Entropy: 50.5, Strength: StrengthInfo{Level: "Strong", CrackTime: "centuries"}},
	}
	for _, format := range []string{"simple", "json", "csv", "table"} {
		out := captureOutput(func() {
			if err := Output(results, format, true); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
		if strings.TrimSpace(out) != "abc123" {
			t.Errorf("quiet mode for %s: expected 'abc123', got %q", format, strings.TrimSpace(out))
		}
	}
}

func TestOutput_InvalidFormat(t *testing.T) {
	results := []PasswordResult{{Password: "abc123"}}
	if err := Output(results, "invalid", false); err != ErrInvalidOutputFormat {
		t.Errorf("expected ErrInvalidOutputFormat, got %v", err)
	}
}

func TestOutputCheck_Simple(t *testing.T) {
	results := []CheckResult{
		{Password: "weak", Entropy: 20.0, Strength: StrengthInfo{Level: "Weak", CrackTime: "< 1 year", Score: 2}, Suggestions: []string{"add length"}},
	}
	out := captureOutput(func() {
		if err := OutputCheck(results, "simple"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	if !strings.Contains(out, "weak") {
		t.Errorf("expected simple check output to contain password")
	}
}

func TestOutputCheck_JSON(t *testing.T) {
	results := []CheckResult{
		{Password: "strong123!", Entropy: 80.0, Strength: StrengthInfo{Level: "Very Strong", CrackTime: "millennia+", Score: 5}},
	}
	out := captureOutput(func() {
		if err := OutputCheck(results, "json"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	var parsed []map[string]any
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("failed to parse JSON: %v", err)
	}
	if len(parsed) != 1 {
		t.Fatalf("expected 1 result, got %d", len(parsed))
	}
}

func TestOutputCheck_CSV(t *testing.T) {
	results := []CheckResult{
		{Password: "test", Entropy: 30.0, Strength: StrengthInfo{Level: "Reasonable", CrackTime: "years", Score: 3}},
	}
	out := captureOutput(func() {
		if err := OutputCheck(results, "csv"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected header + 1 data row, got %d lines", len(lines))
	}
}

func TestOutputCheck_Table(t *testing.T) {
	results := []CheckResult{
		{Password: "test123!", Entropy: 60.0, Strength: StrengthInfo{Level: "Strong", CrackTime: "centuries", Score: 4}},
	}
	out := captureOutput(func() {
		if err := OutputCheck(results, "table"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	if !strings.Contains(out, "test123!") {
		t.Errorf("expected table to contain password")
	}
}

func TestOutputCheck_TableWithSuggestions(t *testing.T) {
	results := []CheckResult{
		{Password: "short", Entropy: 20.0, Strength: StrengthInfo{Level: "Weak", CrackTime: "< 1 year", Score: 2}, Suggestions: []string{"increase length", "add symbols"}},
	}
	out := captureOutput(func() {
		if err := OutputCheck(results, "table"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	if !strings.Contains(out, "Suggestions:") {
		t.Errorf("expected table with suggestions to contain 'Suggestions:'")
	}
}

func TestValidateOutputFormat(t *testing.T) {
	for _, format := range []string{"simple", "json", "csv", "table"} {
		if err := ValidateOutputFormat(format); err != nil {
			t.Errorf("expected %s to be valid, got error: %v", format, err)
		}
	}
	if err := ValidateOutputFormat("invalid"); err != ErrInvalidOutputFormat {
		t.Errorf("expected invalid format to return ErrInvalidOutputFormat, got %v", err)
	}
}

func TestHasSuggestions(t *testing.T) {
	if hasSuggestions([]CheckResult{{Suggestions: []string{}}}) {
		t.Error("expected false for empty suggestions")
	}
	if !hasSuggestions([]CheckResult{{Suggestions: []string{"add length"}}}) {
		t.Error("expected true for non-empty suggestions")
	}
}

func TestJoinSuggestions(t *testing.T) {
	if joinSuggestions([]string{"a", "b"}) != "a; b" {
		t.Errorf("expected 'a; b', got %q", joinSuggestions([]string{"a", "b"}))
	}
	if joinSuggestions([]string{}) != "" {
		t.Errorf("expected empty string, got %q", joinSuggestions([]string{}))
	}
}

func TestOutput_EmptyResults(t *testing.T) {
	out := captureOutput(func() {
		if err := Output([]PasswordResult{}, "simple", false); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	if strings.TrimSpace(out) != "" {
		t.Errorf("expected empty output, got %q", out)
	}
}

func TestOutputCheck_EmptyResults(t *testing.T) {
	out := captureOutput(func() {
		if err := OutputCheck([]CheckResult{}, "simple"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	if strings.TrimSpace(out) != "" {
		t.Errorf("expected empty output, got %q", out)
	}
}

func TestOutput_TableStyleUsesDefault(t *testing.T) {
	_ = table.StyleDefault
	results := []PasswordResult{{Password: "x", Entropy: 10.0, Strength: StrengthInfo{Level: "Weak", CrackTime: "days"}}}
	out := captureOutput(func() {
		if err := Output(results, "table", false); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	if out == "" {
		t.Error("expected non-empty table output")
	}
}
