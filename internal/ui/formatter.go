package ui

import (
	"encoding/json"
	"fmt"
)

type Formatter interface {
	Format(results []PasswordResult, quiet bool) (string, error)
}

type SimpleFormatter struct{}

func (f *SimpleFormatter) Format(results []PasswordResult, quiet bool) (string, error) {
	var output string
	for _, r := range results {
		if quiet {
			output += r.Password + "\n"
		} else {
			output += fmt.Sprintf("Password: %s\n", r.Password)
			if r.Entropy > 0 {
				output += fmt.Sprintf("  Entropy: %.2f bits\n", r.Entropy)
			}
			if r.Strength.Level != "" {
				output += fmt.Sprintf("  Strength: %s\n", r.Strength.Level)
				output += fmt.Sprintf("  Crack Time: %s\n", r.Strength.CrackTime)
			}
			output += "\n"
		}
	}
	return output, nil
}

type JSONFormatter struct{}

func (f *JSONFormatter) Format(results []PasswordResult, quiet bool) (string, error) {
	type output struct {
		Password string  `json:"password"`
		Entropy  float64 `json:"entropy,omitempty"`
		Strength *struct {
			Level     string `json:"level"`
			CrackTime string `json:"crack_time"`
			Score     int    `json:"score"`
		} `json:"strength,omitempty"`
	}

	var outputData []output
	for _, r := range results {
		o := output{Password: r.Password}
		if !quiet && r.Entropy > 0 {
			o.Entropy = r.Entropy
		}
		if !quiet && r.Strength.Level != "" {
			o.Strength = &struct {
				Level     string `json:"level"`
				CrackTime string `json:"crack_time"`
				Score     int    `json:"score"`
			}{
				Level:     r.Strength.Level,
				CrackTime: r.Strength.CrackTime,
				Score:     r.Strength.Score,
			}
		}
		outputData = append(outputData, o)
	}

	data, err := json.MarshalIndent(outputData, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GetFormatter(format string) Formatter {
	switch format {
	case "json":
		return &JSONFormatter{}
	default:
		return &SimpleFormatter{}
	}
}
