package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	PaymentStatementsCmd.AddCommand(PaymentStatementsListCmd)
}

// PaymentStatementsListCmd defines 'list' subcommand
var PaymentStatementsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/payment_statements:get:summary"),
	Long:  TRAPI(`/payment_statements:get:description`),
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

		param, err := collectPaymentStatementsListCmdParams(ac)
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

func collectPaymentStatementsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForPaymentStatementsListCmd("/payment_statements"),
		query:  buildQueryForPaymentStatementsListCmd(),
	}, nil
}

func buildPathForPaymentStatementsListCmd(path string) string {

	return path
}

func buildQueryForPaymentStatementsListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
