package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	UsersCmd.AddCommand(UsersPasswordCmd)
}

// UsersPasswordCmd defines 'password' subcommand
var UsersPasswordCmd = &cobra.Command{
	Use:   "password",
	Short: TRCLI("cli.users.password.summary"),
	Long:  TRCLI(`cli.users.password.description`),
}
