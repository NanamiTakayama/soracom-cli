package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesGetCmdDeviceId holds value of 'device_id' option
var DevicesGetCmdDeviceId string

// DevicesGetCmdModel holds value of 'model' option
var DevicesGetCmdModel bool

func init() {
	DevicesGetCmd.Flags().StringVar(&DevicesGetCmdDeviceId, "device-id", "", TRAPI("Device ID"))

	DevicesGetCmd.Flags().BoolVar(&DevicesGetCmdModel, "model", false, TRAPI("Whether or not to add model information"))

	DevicesCmd.AddCommand(DevicesGetCmd)
}

// DevicesGetCmd defines 'get' subcommand
var DevicesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/devices/{device_id}:get:summary"),
	Long:  TRAPI(`/devices/{device_id}:get:description`),
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

		param, err := collectDevicesGetCmdParams(ac)
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

func collectDevicesGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesGetCmd("/devices/{device_id}"),
		query:  buildQueryForDevicesGetCmd(),
	}, nil
}

func buildPathForDevicesGetCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", DevicesGetCmdDeviceId, -1)

	return path
}

func buildQueryForDevicesGetCmd() string {
	result := []string{}

	if DevicesGetCmdModel != false {
		result = append(result, sprintf("%s=%t", "model", DevicesGetCmdModel))
	}

	return strings.Join(result, "&")
}
