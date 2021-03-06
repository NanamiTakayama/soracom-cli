package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var DevicesListCmdLastEvaluatedKey string

// DevicesListCmdTagName holds value of 'tag_name' option
var DevicesListCmdTagName string

// DevicesListCmdTagValue holds value of 'tag_value' option
var DevicesListCmdTagValue string

// DevicesListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var DevicesListCmdTagValueMatchMode string

// DevicesListCmdLimit holds value of 'limit' option
var DevicesListCmdLimit int64

func init() {
	DevicesListCmd.Flags().StringVar(&DevicesListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("ID of the last Device in the previous page"))

	DevicesListCmd.Flags().StringVar(&DevicesListCmdTagName, "tag-name", "", TRAPI("Tag name"))

	DevicesListCmd.Flags().StringVar(&DevicesListCmdTagValue, "tag-value", "", TRAPI("Tag value"))

	DevicesListCmd.Flags().StringVar(&DevicesListCmdTagValueMatchMode, "tag-value-match-mode", "", TRAPI("Tag value match mode (exact | prefix)"))

	DevicesListCmd.Flags().Int64Var(&DevicesListCmdLimit, "limit", 0, TRAPI("Max number of Devices in a response"))

	DevicesCmd.AddCommand(DevicesListCmd)
}

// DevicesListCmd defines 'list' subcommand
var DevicesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/devices:get:summary"),
	Long:  TRAPI(`/devices:get:description`),
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

		param, err := collectDevicesListCmdParams(ac)
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

func collectDevicesListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesListCmd("/devices"),
		query:  buildQueryForDevicesListCmd(),
	}, nil
}

func buildPathForDevicesListCmd(path string) string {

	return path
}

func buildQueryForDevicesListCmd() string {
	result := []string{}

	if DevicesListCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", DevicesListCmdLastEvaluatedKey))
	}

	if DevicesListCmdTagName != "" {
		result = append(result, sprintf("%s=%s", "tag_name", DevicesListCmdTagName))
	}

	if DevicesListCmdTagValue != "" {
		result = append(result, sprintf("%s=%s", "tag_value", DevicesListCmdTagValue))
	}

	if DevicesListCmdTagValueMatchMode != "" {
		result = append(result, sprintf("%s=%s", "tag_value_match_mode", DevicesListCmdTagValueMatchMode))
	}

	if DevicesListCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", DevicesListCmdLimit))
	}

	return strings.Join(result, "&")
}
