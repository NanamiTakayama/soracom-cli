package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(FilesCmd)
}

// FilesCmd defines 'files' subcommand
var FilesCmd = &cobra.Command{
	Use:   "files",
	Short: TRCLI("cli.files.summary"),
	Long:  TRCLI(`cli.files.description`),
}
