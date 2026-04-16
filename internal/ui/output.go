package ui

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/trust-forge-capital/ohmypassword/internal/i18n"
)

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
	switch format {
	case "json":
		return outputJSON(results, quiet)
	case "csv":
		return outputCSV(results, quiet)
	case "table":
		return outputTable(results, quiet)
	default:
		return outputSimple(results, quiet)
	}
}

func outputSimple(results []PasswordResult, quiet bool) error {
	for _, r := range results {
		if quiet {
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

func outputJSON(results []PasswordResult, quiet bool) error {
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

func outputCSV(results []PasswordResult, quiet bool) error {
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()

	header := []string{"password"}
	if !quiet {
		header = append(header, "entropy", "strength", "crack_time")
	}
	w.Write(header)

	for _, r := range results {
		row := []string{r.Password}
		if !quiet {
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
		w.Write(row)
	}
	return nil
}

func outputTable(results []PasswordResult, quiet bool) error {
	type tableRow struct {
		Password  string
		Entropy   string
		Strength  string
		CrackTime string
	}

	var tableData []tableRow
	for _, r := range results {
		row := tableRow{Password: r.Password}
		if !quiet {
			if r.Entropy > 0 {
				row.Entropy = fmt.Sprintf("%.2f bits", r.Entropy)
			}
			if r.Strength.Level != "" {
				row.Strength = r.Strength.Level
				row.CrackTime = r.Strength.CrackTime
			}
		}
		tableData = append(tableData, row)
	}

	format := "%-30s %-15s %-12s %-20s\n"
	if quiet {
		for _, r := range results {
			fmt.Println(r.Password)
		}
		return nil
	}

	fmt.Printf(format, "PASSWORD", "ENTROPY", "STRENGTH", "CRACK TIME")
	fmt.Println(strings.Repeat("-", 77))
	for _, r := range tableData {
		fmt.Printf(format, r.Password, r.Entropy, r.Strength, r.CrackTime)
	}

	return nil
}
