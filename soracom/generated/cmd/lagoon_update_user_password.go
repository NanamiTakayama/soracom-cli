package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonUpdateUserPasswordCmdNewPassword holds value of 'newPassword' option
var LagoonUpdateUserPasswordCmdNewPassword string

// LagoonUpdateUserPasswordCmdOldPassword holds value of 'oldPassword' option
var LagoonUpdateUserPasswordCmdOldPassword string

// LagoonUpdateUserPasswordCmdUserEmail holds value of 'userEmail' option
var LagoonUpdateUserPasswordCmdUserEmail string

// LagoonUpdateUserPasswordCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUpdateUserPasswordCmdLagoonUserId int64

// LagoonUpdateUserPasswordCmdBody holds contents of request body to be sent
var LagoonUpdateUserPasswordCmdBody string

func init() {
	LagoonUpdateUserPasswordCmd.Flags().StringVar(&LagoonUpdateUserPasswordCmdNewPassword, "new-password", "", TRAPI(""))

	LagoonUpdateUserPasswordCmd.Flags().StringVar(&LagoonUpdateUserPasswordCmdOldPassword, "old-password", "", TRAPI(""))

	LagoonUpdateUserPasswordCmd.Flags().StringVar(&LagoonUpdateUserPasswordCmdUserEmail, "user-email", "", TRAPI(""))

	LagoonUpdateUserPasswordCmd.Flags().Int64Var(&LagoonUpdateUserPasswordCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUpdateUserPasswordCmd.Flags().StringVar(&LagoonUpdateUserPasswordCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LagoonCmd.AddCommand(LagoonUpdateUserPasswordCmd)
}

// LagoonUpdateUserPasswordCmd defines 'update-user-password' subcommand
var LagoonUpdateUserPasswordCmd = &cobra.Command{
	Use:   "update-user-password",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}/password:put:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}/password:put:description`),
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

		param, err := collectLagoonUpdateUserPasswordCmdParams(ac)
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

func collectLagoonUpdateUserPasswordCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonUpdateUserPasswordCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUpdateUserPasswordCmd("/lagoon/users/{lagoon_user_id}/password"),
		query:       buildQueryForLagoonUpdateUserPasswordCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonUpdateUserPasswordCmd(path string) string {

	path = strings.Replace(path, "{"+"lagoon_user_id"+"}", sprintf("%d", LagoonUpdateUserPasswordCmdLagoonUserId), -1)

	return path
}

func buildQueryForLagoonUpdateUserPasswordCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLagoonUpdateUserPasswordCmd() (string, error) {
	if LagoonUpdateUserPasswordCmdBody != "" {
		if strings.HasPrefix(LagoonUpdateUserPasswordCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUpdateUserPasswordCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LagoonUpdateUserPasswordCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LagoonUpdateUserPasswordCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if LagoonUpdateUserPasswordCmdNewPassword != "" {
		result["newPassword"] = LagoonUpdateUserPasswordCmdNewPassword
	}

	if LagoonUpdateUserPasswordCmdOldPassword != "" {
		result["oldPassword"] = LagoonUpdateUserPasswordCmdOldPassword
	}

	if LagoonUpdateUserPasswordCmdUserEmail != "" {
		result["userEmail"] = LagoonUpdateUserPasswordCmdUserEmail
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
