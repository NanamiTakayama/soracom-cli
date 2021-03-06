package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesGetObjectModelCmdModelId holds value of 'model_id' option
var DevicesGetObjectModelCmdModelId string

func init() {
	DevicesGetObjectModelCmd.Flags().StringVar(&DevicesGetObjectModelCmdModelId, "model-id", "", TRAPI("Device object model ID"))

	DevicesCmd.AddCommand(DevicesGetObjectModelCmd)
}

// DevicesGetObjectModelCmd defines 'get-object-model' subcommand
var DevicesGetObjectModelCmd = &cobra.Command{
	Use:   "get-object-model",
	Short: TRAPI("/device_object_models/{model_id}:get:summary"),
	Long:  TRAPI(`/device_object_models/{model_id}:get:description`),
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

		param, err := collectDevicesGetObjectModelCmdParams(ac)
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

func collectDevicesGetObjectModelCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesGetObjectModelCmd("/device_object_models/{model_id}"),
		query:  buildQueryForDevicesGetObjectModelCmd(),
	}, nil
}

func buildPathForDevicesGetObjectModelCmd(path string) string {

	path = strings.Replace(path, "{"+"model_id"+"}", DevicesGetObjectModelCmdModelId, -1)

	return path
}

func buildQueryForDevicesGetObjectModelCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
