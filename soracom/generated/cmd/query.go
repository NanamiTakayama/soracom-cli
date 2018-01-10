package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(QueryCmd)
}

// QueryCmd defines 'query' subcommand
var QueryCmd = &cobra.Command{
	Use:   "query",
	Short: TRCLI("cli.query.summary"),
	Long:  TRCLI(`cli.query.description`),
}