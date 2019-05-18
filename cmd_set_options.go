package main

import (
	"github.com/spf13/cobra"
)

func createCmdSetOptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setoptions",
		Short: "Set options on an account.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(
		createCmdSetOptionsInflationDestination(),
	)
	return cmd
}
