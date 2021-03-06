package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorEnableMfaCmdOperatorId holds value of 'operator_id' option
var OperatorEnableMfaCmdOperatorId string

func init() {
	OperatorEnableMfaCmd.Flags().StringVar(&OperatorEnableMfaCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorCmd.AddCommand(OperatorEnableMfaCmd)
}

// OperatorEnableMfaCmd defines 'enable-mfa' subcommand
var OperatorEnableMfaCmd = &cobra.Command{
	Use:   "enable-mfa",
	Short: TRAPI("/operators/{operator_id}/mfa:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/mfa:post:description`),
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

		param, err := collectOperatorEnableMfaCmdParams(ac)
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

func collectOperatorEnableMfaCmdParams(ac *apiClient) (*apiParams, error) {

	if OperatorEnableMfaCmdOperatorId == "" {
		OperatorEnableMfaCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForOperatorEnableMfaCmd("/operators/{operator_id}/mfa"),
		query:  buildQueryForOperatorEnableMfaCmd(),
	}, nil
}

func buildPathForOperatorEnableMfaCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorEnableMfaCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorEnableMfaCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
