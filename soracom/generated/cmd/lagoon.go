package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(LagoonCmd)
}

// LagoonCmd defines 'lagoon' subcommand
var LagoonCmd = &cobra.Command{
	Use:   "lagoon",
	Short: TRCLI("cli.lagoon.summary"),
	Long:  TRCLI(`cli.lagoon.description`),
}
