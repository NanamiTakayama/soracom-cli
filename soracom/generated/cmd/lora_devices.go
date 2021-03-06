package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(LoraDevicesCmd)
}

// LoraDevicesCmd defines 'lora-devices' subcommand
var LoraDevicesCmd = &cobra.Command{
	Use:   "lora-devices",
	Short: TRCLI("cli.lora-devices.summary"),
	Long:  TRCLI(`cli.lora-devices.description`),
}
