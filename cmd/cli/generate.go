package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trust-forge-capital/ohmypassword/internal/generator"
	"github.com/trust-forge-capital/ohmypassword/internal/i18n"
	"github.com/trust-forge-capital/ohmypassword/internal/ui"
	"github.com/trust-forge-capital/ohmypassword/internal/validator"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   i18n.T("generate_use"),
	Long:    i18n.T("generate_long"),
	Aliases: []string{"gen"},
	RunE: func(cmd *cobra.Command, args []string) error {
		length, _ := cmd.Flags().GetInt("length")
		charset, _ := cmd.Flags().GetString("charset")
		strategy, _ := cmd.Flags().GetString("strategy")
		count, _ := cmd.Flags().GetInt("count")
		validate, _ := cmd.Flags().GetBool("validate")
		quiet, _ := cmd.Flags().GetBool("quiet")
		output, _ := cmd.Flags().GetString("output")
		excludeSimilar, _ := cmd.Flags().GetBool("exclude-similar")

		opts := &generator.Options{
			Length:         length,
			Charset:        charset,
			Strategy:       strategy,
			Count:          count,
			Validate:       validate,
			Quiet:          quiet,
			ExcludeSimilar: excludeSimilar,
		}

		passwords, err := generator.GeneratePasswords(opts)
		if err != nil {
			return err
		}

		results := make([]ui.PasswordResult, len(passwords))
		for i, pwd := range passwords {
			result := ui.PasswordResult{Password: pwd}
			if validate {
				strength := validator.CalculateStrength(pwd, opts.Charset)
				result.Strength = ui.StrengthInfo{
					Level:     validator.GetDisplayName(validator.StrengthLevel(strength.Level)),
					CrackTime: strength.CrackTime,
					Score:     strength.Score,
				}
			}
			results[i] = result
		}

		return ui.Output(results, output, quiet)
	},
}

func init() {
	generateCmd.Flags().IntP("length", "l", 16, i18n.T("flag_length"))
	generateCmd.Flags().StringP("charset", "c", "all", i18n.T("flag_charset"))
	generateCmd.Flags().StringP("strategy", "s", "simple", i18n.T("flag_strategy"))
	generateCmd.Flags().IntP("count", "n", 1, i18n.T("flag_count"))
	generateCmd.Flags().BoolP("validate", "v", false, i18n.T("flag_validate"))
	generateCmd.Flags().BoolP("quiet", "q", false, i18n.T("flag_quiet"))
	generateCmd.Flags().Bool("exclude-similar", false, i18n.T("flag_exclude_similar"))

	RootCmd.AddCommand(generateCmd)
}
