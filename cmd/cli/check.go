package cli

import (
	"github.com/spf13/cobra"
	"github.com/trust-forge-capital/ohmypassword/internal/i18n"
	"github.com/trust-forge-capital/ohmypassword/internal/ui"
	"github.com/trust-forge-capital/ohmypassword/internal/validator"
)

var CheckCmd = &cobra.Command{
	Use:     "check",
	Short:   i18n.T("check_use"),
	Long:    i18n.T("check_long"),
	Aliases: []string{"ck"},
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		results := make([]ui.CheckResult, len(args))
		for i, password := range args {
			strength := validator.CalculateStrength(password, "all")
			results[i] = ui.CheckResult{
				Password: password,
				Entropy:  strength.Entropy,
				Strength: ui.StrengthInfo{
					Level:     validator.GetDisplayName(validator.StrengthLevel(strength.Level)),
					CrackTime: strength.CrackTime,
					Score:     strength.Score,
				},
				Suggestions: strength.Suggestions,
			}
		}

		err = ui.OutputCheck(results, output)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	CheckCmd.Flags().StringP("output", "o", "simple", i18n.T("flag_output"))
}
