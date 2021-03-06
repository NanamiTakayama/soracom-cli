package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersUnsetExpiryTimeCmdImsi holds value of 'imsi' option
var SubscribersUnsetExpiryTimeCmdImsi string

func init() {
	SubscribersUnsetExpiryTimeCmd.Flags().StringVar(&SubscribersUnsetExpiryTimeCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersCmd.AddCommand(SubscribersUnsetExpiryTimeCmd)
}

// SubscribersUnsetExpiryTimeCmd defines 'unset-expiry-time' subcommand
var SubscribersUnsetExpiryTimeCmd = &cobra.Command{
	Use:   "unset-expiry-time",
	Short: TRAPI("/subscribers/{imsi}/unset_expiry_time:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/unset_expiry_time:post:description`),
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

		param, err := collectSubscribersUnsetExpiryTimeCmdParams(ac)
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

func collectSubscribersUnsetExpiryTimeCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersUnsetExpiryTimeCmd("/subscribers/{imsi}/unset_expiry_time"),
		query:  buildQueryForSubscribersUnsetExpiryTimeCmd(),
	}, nil
}

func buildPathForSubscribersUnsetExpiryTimeCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersUnsetExpiryTimeCmdImsi, -1)

	return path
}

func buildQueryForSubscribersUnsetExpiryTimeCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
