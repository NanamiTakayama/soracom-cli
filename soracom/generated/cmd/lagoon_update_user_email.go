package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonUpdateUserEmailCmdUserEmail holds value of 'userEmail' option
var LagoonUpdateUserEmailCmdUserEmail string

// LagoonUpdateUserEmailCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUpdateUserEmailCmdLagoonUserId int64

// LagoonUpdateUserEmailCmdBody holds contents of request body to be sent
var LagoonUpdateUserEmailCmdBody string

func init() {
	LagoonUpdateUserEmailCmd.Flags().StringVar(&LagoonUpdateUserEmailCmdUserEmail, "user-email", "", TRAPI(""))

	LagoonUpdateUserEmailCmd.Flags().Int64Var(&LagoonUpdateUserEmailCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUpdateUserEmailCmd.Flags().StringVar(&LagoonUpdateUserEmailCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LagoonCmd.AddCommand(LagoonUpdateUserEmailCmd)
}

// LagoonUpdateUserEmailCmd defines 'update-user-email' subcommand
var LagoonUpdateUserEmailCmd = &cobra.Command{
	Use:   "update-user-email",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}/email:put:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}/email:put:description`),
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

		param, err := collectLagoonUpdateUserEmailCmdParams(ac)
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

func collectLagoonUpdateUserEmailCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonUpdateUserEmailCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUpdateUserEmailCmd("/lagoon/users/{lagoon_user_id}/email"),
		query:       buildQueryForLagoonUpdateUserEmailCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonUpdateUserEmailCmd(path string) string {

	path = strings.Replace(path, "{"+"lagoon_user_id"+"}", sprintf("%d", LagoonUpdateUserEmailCmdLagoonUserId), -1)

	return path
}

func buildQueryForLagoonUpdateUserEmailCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLagoonUpdateUserEmailCmd() (string, error) {
	if LagoonUpdateUserEmailCmdBody != "" {
		if strings.HasPrefix(LagoonUpdateUserEmailCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUpdateUserEmailCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LagoonUpdateUserEmailCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LagoonUpdateUserEmailCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if LagoonUpdateUserEmailCmdUserEmail != "" {
		result["userEmail"] = LagoonUpdateUserEmailCmdUserEmail
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
