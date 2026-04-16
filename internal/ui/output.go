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

type CheckResult struct {
	Password    string
	Entropy     float64
	Strength    StrengthInfo
	Suggestions []string
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

func OutputCheck(results []CheckResult, format string) error {
	if err := ValidateOutputFormat(format); err != nil {
		return err
	}

	switch format {
	case "json":
		return outputCheckJSON(results)
	case "csv":
		return outputCheckCSV(results)
	case "table":
		return outputCheckTable(results)
	default:
		return outputCheckSimple(results)
	}
}

func outputCheckSimple(results []CheckResult) error {
	for _, r := range results {
		fmt.Printf("Password: %s\n", r.Password)
		fmt.Printf("  Score: %d/5\n", r.Strength.Score)
		fmt.Printf("  %s: %.2f bits\n", i18n.T("output_entropy"), r.Entropy)
		fmt.Printf("  %s: %s\n", i18n.T("output_strength"), r.Strength.Level)
		fmt.Printf("  %s: %s\n", i18n.T("output_crack_time"), r.Strength.CrackTime)
		if len(r.Suggestions) > 0 {
			fmt.Println("  Suggestions:")
			for _, s := range r.Suggestions {
				fmt.Printf("    - %s\n", s)
			}
		}
		fmt.Println()
	}
	return nil
}

type checkJSONOutput struct {
	Password    string   `json:"password"`
	Score       int      `json:"score"`
	Entropy     float64  `json:"entropy"`
	Strength    string   `json:"strength"`
	CrackTime   string   `json:"crack_time"`
	Suggestions []string `json:"suggestions,omitempty"`
}

func outputCheckJSON(results []CheckResult) error {
	var outputData []checkJSONOutput
	for _, r := range results {
		o := checkJSONOutput{
			Password:    r.Password,
			Score:       r.Strength.Score,
			Entropy:     r.Entropy,
			Strength:    r.Strength.Level,
			CrackTime:   r.Strength.CrackTime,
			Suggestions: r.Suggestions,
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

func outputCheckCSV(results []CheckResult) error {
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()

	header := []string{"password", "score", "entropy", "strength", "crack_time"}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, r := range results {
		row := []string{
			r.Password,
			fmt.Sprintf("%d/5", r.Strength.Score),
			fmt.Sprintf("%.2f", r.Entropy),
			r.Strength.Level,
			r.Strength.CrackTime,
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func outputCheckTable(results []CheckResult) error {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleDefault)

	t.AppendHeader(table.Row{
		i18n.T("output_password"),
		"SCORE",
		i18n.T("output_entropy"),
		i18n.T("output_strength"),
		i18n.T("output_crack_time"),
	})

	for _, r := range results {
		row := table.Row{
			r.Password,
			fmt.Sprintf("%d/5", r.Strength.Score),
			fmt.Sprintf("%.2f bits", r.Entropy),
			r.Strength.Level,
			r.Strength.CrackTime,
		}
		t.AppendRow(row)
	}

	t.Render()

	if hasSuggestions(results) {
		fmt.Println("\nSuggestions:")
		t2 := table.NewWriter()
		t2.SetOutputMirror(os.Stdout)
		t2.SetStyle(table.StyleDefault)
		t2.AppendHeader(table.Row{"PASSWORD", "SUGGESTIONS"})
		for _, r := range results {
			if len(r.Suggestions) > 0 {
				t2.AppendRow(table.Row{r.Password, joinSuggestions(r.Suggestions)})
			}
		}
		t2.Render()
	}

	return nil
}

func hasSuggestions(results []CheckResult) bool {
	for _, r := range results {
		if len(r.Suggestions) > 0 {
			return true
		}
	}
	return false
}

func joinSuggestions(s []string) string {
	result := ""
	for i, s := range s {
		if i > 0 {
			result += "; "
		}
		result += s
	}
	return result
}
