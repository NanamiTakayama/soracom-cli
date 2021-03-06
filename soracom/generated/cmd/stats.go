package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(StatsCmd)
}

// StatsCmd defines 'stats' subcommand
var StatsCmd = &cobra.Command{
	Use:   "stats",
	Short: TRCLI("cli.stats.summary"),
	Long:  TRCLI(`cli.stats.description`),
}
