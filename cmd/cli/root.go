package cli

import (
	"fmt"
	"os"
	"strings"

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
	Short: "ohmypassword",
	Long:  "Secure password generator",
}

func init() {
	lang := ""
	for i, arg := range os.Args {
		if (arg == "-L" || arg == "--lang") && i+1 < len(os.Args) {
			lang = os.Args[i+1]
			break
		}
	}

	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		if firstArg != "generate" && firstArg != "gen" && firstArg != "check" && firstArg != "ck" && firstArg != "version" && firstArg != "completion" && firstArg != "help" && firstArg != "-h" && firstArg != "--help" && firstArg != "-v" && firstArg != "--version" {
			newArgs := make([]string, 0, len(os.Args)+1)
			newArgs = append(newArgs, os.Args[0])
			newArgs = append(newArgs, "generate")
			newArgs = append(newArgs, os.Args[1:]...)
			os.Args = newArgs

			if lang == "" {
				for i, arg := range os.Args {
					if (arg == "-L" || arg == "--lang") && i+1 < len(os.Args) {
						lang = os.Args[i+1]
						break
					}
				}
			}
		}
	} else if len(os.Args) == 1 {
		newArgs := []string{os.Args[0], "generate"}
		os.Args = newArgs
	}

	if lang != "" {
		i18n.SetLanguage(lang)
	}

	RootCmd.AddCommand(GenerateCmd)
	RootCmd.AddCommand(CheckCmd)
	RootCmd.AddCommand(versionCmd)

	RootCmd.PersistentFlags().StringP("lang", "L", "", i18n.T("flag_lang"))
	RootCmd.PersistentFlags().StringP("output", "o", "simple", i18n.T("flag_output"))

	updateCommandStrings()

	versionCmd.Version = fmt.Sprintf("%s (%s %s)", version, gitCommit, buildTime)
}

func updateCommandStrings() {
	RootCmd.Short = i18n.T("root_use")
	RootCmd.Long = i18n.T("root_long")
	GenerateCmd.Short = i18n.T("generate_use")
	GenerateCmd.Long = i18n.T("generate_long")
	CheckCmd.Short = i18n.T("check_use")
	CheckCmd.Long = i18n.T("check_long")
	versionCmd.Short = i18n.T("version_use")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: i18n.T("version_use"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version)
		fmt.Println("Commit:", gitCommit)
		fmt.Println("Build:", strings.ReplaceAll(buildTime, "_", " "))
	},
}
