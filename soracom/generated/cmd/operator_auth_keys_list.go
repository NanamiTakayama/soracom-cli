package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorAuthKeysListCmdOperatorId holds value of 'operator_id' option
var OperatorAuthKeysListCmdOperatorId string

func init() {
	OperatorAuthKeysListCmd.Flags().StringVar(&OperatorAuthKeysListCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorAuthKeysCmd.AddCommand(OperatorAuthKeysListCmd)
}

// OperatorAuthKeysListCmd defines 'list' subcommand
var OperatorAuthKeysListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/auth_keys:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/auth_keys:get:description`),
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

		param, err := collectOperatorAuthKeysListCmdParams(ac)
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

func collectOperatorAuthKeysListCmdParams(ac *apiClient) (*apiParams, error) {

	if OperatorAuthKeysListCmdOperatorId == "" {
		OperatorAuthKeysListCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForOperatorAuthKeysListCmd("/operators/{operator_id}/auth_keys"),
		query:  buildQueryForOperatorAuthKeysListCmd(),
	}, nil
}

func buildPathForOperatorAuthKeysListCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorAuthKeysListCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorAuthKeysListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
