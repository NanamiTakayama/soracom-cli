package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	LagoonCmd.AddCommand(LagoonListUsersCmd)
}

// LagoonListUsersCmd defines 'list-users' subcommand
var LagoonListUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: TRAPI("/lagoon/users:get:summary"),
	Long:  TRAPI(`/lagoon/users:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectLagoonListUsersCmdParams(ac)
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectLagoonListUsersCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLagoonListUsersCmd("/lagoon/users"),
		query:  buildQueryForLagoonListUsersCmd(),
	}, nil
}

func buildPathForLagoonListUsersCmd(path string) string {

	return path
}

func buildQueryForLagoonListUsersCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
