package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// EventHandlersIgnoreCmdHandlerId holds value of 'handler_id' option
var EventHandlersIgnoreCmdHandlerId string

// EventHandlersIgnoreCmdImsi holds value of 'imsi' option
var EventHandlersIgnoreCmdImsi string

func init() {
	EventHandlersIgnoreCmd.Flags().StringVar(&EventHandlersIgnoreCmdHandlerId, "handler-id", "", TRAPI("handler_id"))

	EventHandlersIgnoreCmd.Flags().StringVar(&EventHandlersIgnoreCmdImsi, "imsi", "", TRAPI("imsi"))

	EventHandlersCmd.AddCommand(EventHandlersIgnoreCmd)
}

// EventHandlersIgnoreCmd defines 'ignore' subcommand
var EventHandlersIgnoreCmd = &cobra.Command{
	Use:   "ignore",
	Short: TRAPI("/event_handlers/{handler_id}/subscribers/{imsi}/ignore:post:summary"),
	Long:  TRAPI(`/event_handlers/{handler_id}/subscribers/{imsi}/ignore:post:description`),
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

		param, err := collectEventHandlersIgnoreCmdParams(ac)
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

func collectEventHandlersIgnoreCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForEventHandlersIgnoreCmd("/event_handlers/{handler_id}/subscribers/{imsi}/ignore"),
		query:  buildQueryForEventHandlersIgnoreCmd(),
	}, nil
}

func buildPathForEventHandlersIgnoreCmd(path string) string {

	path = strings.Replace(path, "{"+"handler_id"+"}", EventHandlersIgnoreCmdHandlerId, -1)

	path = strings.Replace(path, "{"+"imsi"+"}", EventHandlersIgnoreCmdImsi, -1)

	return path
}

func buildQueryForEventHandlersIgnoreCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
