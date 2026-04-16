package main

import (
	"os"

	"github.com/trust-forge-capital/ohmypassword/cmd/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}