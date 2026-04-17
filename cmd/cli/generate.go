package cli

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/trust-forge-capital/ohmypassword/internal/generator"
	"github.com/trust-forge-capital/ohmypassword/internal/i18n"
	"github.com/trust-forge-capital/ohmypassword/internal/ui"
	"github.com/trust-forge-capital/ohmypassword/internal/validator"
)

var GenerateCmd = &cobra.Command{
	Use:     "generate",
	Short:   i18n.T("generate_use"),
	Long:    i18n.T("generate_long"),
	Aliases: []string{"gen"},
	RunE: func(cmd *cobra.Command, args []string) error {
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			return err
		}
		charset, err := cmd.Flags().GetString("charset")
		if err != nil {
			return err
		}
		strategy, err := cmd.Flags().GetString("strategy")
		if err != nil {
			return err
		}
		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			return err
		}
		validate, err := cmd.Flags().GetBool("validate")
		if err != nil {
			return err
		}
		quiet, err := cmd.Flags().GetBool("quiet")
		if err != nil {
			return err
		}
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}
		excludeSimilar, err := cmd.Flags().GetBool("exclude-similar")
		if err != nil {
			return err
		}

		if strategy == "passphrase" && !cmd.Flags().Changed("length") {
			length = 4
		}
		if strategy == "segmented" && !cmd.Flags().Changed("length") {
			length = 12
		}

		opts := &generator.Options{
			Length:         length,
			Charset:        charset,
			Strategy:       strategy,
			Count:          count,
			ShowStrength:   validate,
			Quiet:          quiet,
			ExcludeSimilar: excludeSimilar,
		}

		passwords, err := generator.GeneratePasswords(opts)
		if err != nil {
			return translateError(err, strategy)
		}

		results := make([]ui.PasswordResult, len(passwords))
		for i, pwd := range passwords {
			result := ui.PasswordResult{Password: pwd}
			if validate {
				strength := validator.CalculateStrength(pwd, opts.Charset)
				result.Entropy = strength.Entropy
				result.Strength = ui.StrengthInfo{
					Level:     validator.GetDisplayName(validator.StrengthLevel(strength.Level)),
					CrackTime: strength.CrackTime,
					Score:     strength.Score,
				}
			}
			results[i] = result
		}

		err = ui.Output(results, output, quiet)
		if err != nil {
			return translateError(err, strategy)
		}
		return nil
	},
}

func translateError(err error, strategy string) error {
	if errors.Is(err, generator.ErrInvalidLength) {
		if strategy == "passphrase" {
			return errors.New(i18n.T("error_invalid_passphrase_length"))
		}
		return errors.New(i18n.T("error_invalid_length"))
	}
	if errors.Is(err, generator.ErrInvalidCount) {
		return errors.New(i18n.T("error_invalid_count"))
	}
	if errors.Is(err, generator.ErrInvalidStrategy) {
		return errors.New(i18n.T("error_invalid_strategy"))
	}
	if errors.Is(err, generator.ErrInvalidCharset) {
		return errors.New(i18n.T("error_invalid_charset"))
	}
	if errors.Is(err, ui.ErrInvalidOutputFormat) {
		return errors.New(i18n.T("error_invalid_output"))
	}
	return err
}

func init() {
	GenerateCmd.Flags().IntP("length", "l", 16, i18n.T("flag_length"))
	GenerateCmd.Flags().StringP("charset", "c", "all", i18n.T("flag_charset"))
	GenerateCmd.Flags().StringP("strategy", "s", "simple", i18n.T("flag_strategy"))
	GenerateCmd.Flags().IntP("count", "n", 1, i18n.T("flag_count"))
	GenerateCmd.Flags().BoolP("validate", "V", false, i18n.T("flag_validate"))
	GenerateCmd.Flags().BoolP("quiet", "q", false, i18n.T("flag_quiet"))
	GenerateCmd.Flags().Bool("exclude-similar", false, i18n.T("flag_exclude_similar"))
}
