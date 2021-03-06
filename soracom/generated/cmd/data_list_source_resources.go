package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DataListSourceResourcesCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var DataListSourceResourcesCmdLastEvaluatedKey string

// DataListSourceResourcesCmdResourceType holds value of 'resource_type' option
var DataListSourceResourcesCmdResourceType string

// DataListSourceResourcesCmdLimit holds value of 'limit' option
var DataListSourceResourcesCmdLimit int64

func init() {
	DataListSourceResourcesCmd.Flags().StringVar(&DataListSourceResourcesCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of `resourceId` in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	DataListSourceResourcesCmd.Flags().StringVar(&DataListSourceResourcesCmdResourceType, "resource-type", "", TRAPI("Type of data source resource"))

	DataListSourceResourcesCmd.Flags().Int64Var(&DataListSourceResourcesCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	DataCmd.AddCommand(DataListSourceResourcesCmd)
}

// DataListSourceResourcesCmd defines 'list-source-resources' subcommand
var DataListSourceResourcesCmd = &cobra.Command{
	Use:   "list-source-resources",
	Short: TRAPI("/data/resources:get:summary"),
	Long:  TRAPI(`/data/resources:get:description`),
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

		param, err := collectDataListSourceResourcesCmdParams(ac)
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

func collectDataListSourceResourcesCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDataListSourceResourcesCmd("/data/resources"),
		query:  buildQueryForDataListSourceResourcesCmd(),
	}, nil
}

func buildPathForDataListSourceResourcesCmd(path string) string {

	return path
}

func buildQueryForDataListSourceResourcesCmd() string {
	result := []string{}

	if DataListSourceResourcesCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", DataListSourceResourcesCmdLastEvaluatedKey))
	}

	if DataListSourceResourcesCmdResourceType != "" {
		result = append(result, sprintf("%s=%s", "resource_type", DataListSourceResourcesCmdResourceType))
	}

	if DataListSourceResourcesCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", DataListSourceResourcesCmdLimit))
	}

	return strings.Join(result, "&")
}
