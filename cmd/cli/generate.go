package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trust-forge-capital/ohmypassword/internal/generator"
	"github.com/trust-forge-capital/ohmypassword/internal/i18n"
	"github.com/trust-forge-capital/ohmypassword/internal/ui"
	"github.com/trust-forge-capital/ohmypassword/internal/validator"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: i18n.T("generate_use"),
	Long:  i18n.T("generate_long"),
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
				result.Strength = validator.CalculateStrength(pwd, opts.Charset)
			}
			results[i] = result
		}

		return ui.Output(results, output, quiet)
	},
}

func init() {
	GenerateCmd.Flags().IntP("length", "l", 16, i18n.T("flag_length"))
	GenerateCmd.Flags().StringP("charset", "c", "all", i18n.T("flag_charset"))
	GenerateCmd.Flags().StringP("strategy", "s", "simple", i18n.T("flag_strategy"))
	GenerateCmd.Flags().IntP("count", "n", 1, i18n.T("flag_count"))
	GenerateCmd.Flags().BoolP("validate", "v", false, i18n.T("flag_validate"))
	GenerateCmd.Flags().BoolP("quiet", "q", false, i18n.T("flag_quiet"))
	GenerateCmd.Flags().Bool("exclude-similar", false, i18n.T("flag_exclude_similar"))
}