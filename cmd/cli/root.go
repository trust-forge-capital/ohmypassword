package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trust-forge-capital/ohmypassword/internal/i18n"
)

var (
	version   = "1.0.0"
	buildTime = "unknown"
	gitCommit = "unknown"
)

var RootCmd = &cobra.Command{
	Use:   "ohmypassword",
	Short: i18n.T("root_use"),
	Long:  i18n.T("root_long"),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		i18n.SetLanguage(lang)
	},
}

func init() {
	RootCmd.AddCommand(GenerateCmd)
	RootCmd.AddCommand(versionCmd)

	RootCmd.PersistentFlags().StringP("lang", "L", "", i18n.T("flag_lang"))
	RootCmd.PersistentFlags().StringP("output", "o", "simple", i18n.T("flag_output"))

	versionCmd.Version = fmt.Sprintf("%s (%s %s)", version, gitCommit, buildTime)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: i18n.T("version_use"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Version)
	},
}