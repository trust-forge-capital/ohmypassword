package ui

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/trust-forge-capital/ohmypassword/internal/i18n"
)

var ErrInvalidOutputFormat = errors.New("invalid output format")

type PasswordResult struct {
	Password string
	Entropy  float64
	Strength StrengthInfo
}

type StrengthInfo struct {
	Level     string
	CrackTime string
	Score     int
}

func Output(results []PasswordResult, format string, quiet bool) error {
	if err := ValidateOutputFormat(format); err != nil {
		return err
	}

	hasDetails := false
	for _, r := range results {
		if r.Entropy > 0 || r.Strength.Level != "" {
			hasDetails = true
			break
		}
	}

	switch format {
	case "json":
		return outputJSON(results, quiet, hasDetails)
	case "csv":
		return outputCSV(results, quiet, hasDetails)
	case "table":
		return outputTable(results, quiet, hasDetails)
	default:
		return outputSimple(results, quiet, hasDetails)
	}
}

func ValidateOutputFormat(format string) error {
	switch format {
	case "simple", "json", "csv", "table":
		return nil
	default:
		return ErrInvalidOutputFormat
	}
}

func outputSimple(results []PasswordResult, quiet bool, hasDetails bool) error {
	for _, r := range results {
		if quiet || !hasDetails {
			fmt.Println(r.Password)
		} else {
			fmt.Printf("Password: %s\n", r.Password)
			if r.Entropy > 0 {
				fmt.Printf("  %s: %.2f bits\n", i18n.T("output_entropy"), r.Entropy)
			}
			if r.Strength.Level != "" {
				fmt.Printf("  %s: %s\n", i18n.T("output_strength"), r.Strength.Level)
				fmt.Printf("  %s: %s\n", i18n.T("output_crack_time"), r.Strength.CrackTime)
			}
			fmt.Println()
		}
	}
	return nil
}

func outputJSON(results []PasswordResult, quiet bool, hasDetails bool) error {
	if quiet {
		for _, r := range results {
			fmt.Println(r.Password)
		}
		return nil
	}

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
		return err
	}
	fmt.Println(string(data))
	return nil
}

func outputCSV(results []PasswordResult, quiet bool, hasDetails bool) error {
	if quiet {
		for _, r := range results {
			fmt.Println(r.Password)
		}
		return nil
	}

	w := csv.NewWriter(os.Stdout)
	defer w.Flush()

	header := []string{"password"}
	if !quiet && hasDetails {
		header = append(header, "entropy", "strength", "crack_time")
	}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, r := range results {
		row := []string{r.Password}
		if !quiet && hasDetails {
			if r.Entropy > 0 {
				row = append(row, fmt.Sprintf("%.2f", r.Entropy))
			} else {
				row = append(row, "")
			}
			if r.Strength.Level != "" {
				row = append(row, r.Strength.Level, r.Strength.CrackTime)
			} else {
				row = append(row, "", "")
			}
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func outputTable(results []PasswordResult, quiet bool, hasDetails bool) error {
	if quiet || !hasDetails {
		for _, r := range results {
			fmt.Println(r.Password)
		}
		return nil
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleDefault)

	t.AppendHeader(table.Row{
		i18n.T("output_password"),
		i18n.T("output_entropy"),
		i18n.T("output_strength"),
		i18n.T("output_crack_time"),
	})

	for _, r := range results {
		row := table.Row{r.Password}
		if r.Entropy > 0 {
			row = append(row, fmt.Sprintf("%.2f bits", r.Entropy))
		} else {
			row = append(row, "")
		}
		if r.Strength.Level != "" {
			row = append(row, r.Strength.Level, r.Strength.CrackTime)
		} else {
			row = append(row, "", "")
		}
		t.AppendRow(row)
	}

	t.Render()

	return nil
}
