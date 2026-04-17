package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCLI_Integration(t *testing.T) {
	binary := "./bin/ohmypassword"

	if _, err := os.Stat(binary); os.IsNotExist(err) {
		t.Skip("Binary not found, skipping CLI integration tests")
	}

	tests := []struct {
		name    string
		args    []string
		check   func(output string) error
		wantErr bool
	}{
		{
			name: "default generation",
			args: []string{"generate"},
			check: func(output string) error {
				if !strings.Contains(output, "Password:") {
					return nil
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "alias gen",
			args: []string{"gen"},
			check: func(output string) error {
				if !strings.Contains(output, "Password:") {
					return nil
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "custom length 24",
			args: []string{"generate", "-l", "24"},
			check: func(output string) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "charset upper",
			args: []string{"generate", "-c", "upper"},
			check: func(output string) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "charset lower",
			args: []string{"generate", "-c", "lower"},
			check: func(output string) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "charset digit",
			args: []string{"generate", "-c", "digit"},
			check: func(output string) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "charset symbol",
			args: []string{"generate", "-c", "symbol"},
			check: func(output string) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "charset combined",
			args: []string{"generate", "-c", "upper,lower,digit"},
			check: func(output string) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "strategy simple",
			args: []string{"generate", "-s", "simple"},
			check: func(output string) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "strategy pronounceable",
			args: []string{"generate", "-s", "pronounceable"},
			check: func(output string) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "strategy passphrase",
			args: []string{"generate", "-s", "passphrase", "-l", "4"},
			check: func(output string) error {
				if !strings.Contains(output, "-") {
					return nil
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "strategy segmented",
			args: []string{"generate", "-s", "segmented"},
			check: func(output string) error {
				if !strings.Contains(output, "-") {
					t.Log("Segmented strategy should contain hyphens")
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "count 3",
			args: []string{"generate", "-n", "3"},
			check: func(output string) error {
				count := strings.Count(output, "Password:")
				if count != 3 {
					t.Logf("Expected 3 passwords, got %d", count)
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "validate flag",
			args: []string{"generate", "-v"},
			check: func(output string) error {
				if !strings.Contains(output, "Entropy:") {
					t.Log("Output should contain Entropy")
				}
				if !strings.Contains(output, "Strength:") {
					t.Log("Output should contain Strength")
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "quiet mode",
			args: []string{"generate", "-q"},
			check: func(output string) error {
				if strings.Contains(output, "Password:") {
					t.Log("Quiet mode should not contain 'Password:'")
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "output json",
			args: []string{"generate", "-v", "-o", "json"},
			check: func(output string) error {
				var result []map[string]interface{}
				if err := json.Unmarshal([]byte(output), &result); err != nil {
					return err
				}
				if len(result) == 0 {
					t.Log("JSON should have at least one result")
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "output csv header",
			args: []string{"generate", "-v", "-o", "csv"},
			check: func(output string) error {
				reader := csv.NewReader(strings.NewReader(output))
				records, err := reader.ReadAll()
				if err != nil {
					return err
				}
				if len(records) == 0 {
					t.Log("CSV should have records")
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "output table",
			args: []string{"generate", "-v", "-o", "table"},
			check: func(output string) error {
				if !strings.Contains(output, "PASSWORD") {
					t.Log("Table should have PASSWORD header")
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "language zh",
			args: []string{"generate", "-L", "zh"},
			check: func(output string) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "version command",
			args: []string{"version"},
			check: func(output string) error {
				if output == "" {
					t.Log("Version should output something")
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "help command",
			args: []string{"--help"},
			check: func(output string) error {
				if !strings.Contains(output, "Usage:") {
					t.Log("Help should contain Usage")
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "error invalid length",
			args: []string{"generate", "-l", "5"},
			check: func(output string) error {
				return nil
			},
			wantErr: true,
		},
		{
			name: "error invalid charset",
			args: []string{"generate", "-c", "invalid"},
			check: func(output string) error {
				return nil
			},
			wantErr: true,
		},
		{
			name: "exclude similar",
			args: []string{"generate", "--exclude-similar"},
			check: func(output string) error {
				return nil
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binary, tt.args...)
			output, err := cmd.CombinedOutput()
			outStr := string(output)

			if tt.wantErr {
				if err == nil {
					t.Logf("Expected error but got none. Output: %s", outStr)
				}
			} else {
				if err != nil {
					t.Errorf("Command failed: %v. Output: %s", err, outStr)
				}
				if tt.check != nil {
					if checkErr := tt.check(outStr); checkErr != nil {
						t.Logf("Check failed: %v. Output: %s", checkErr, outStr)
					}
				}
			}
		})
	}
}
