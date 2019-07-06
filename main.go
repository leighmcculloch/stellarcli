package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/network"
)

var test = false
var client *horizonclient.Client
var networkPassphrase string

func main() {
	var cmd = cobra.Command{
		Use:   "stellarscripts",
		Short: "My personal stellar scripts.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmdTestNet := cobra.Command{
		Use:   "test",
		Short: "Run scripts on Stellar testnet.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			test = true
			client = horizonclient.DefaultTestNetClient
			networkPassphrase = network.TestNetworkPassphrase
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmdTestNet.AddCommand(
		createCmdSign(),
	)

	cmdPublicNet := cobra.Command{
		Use:   "public",
		Short: "Run scripts on Stellar publicnet.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			client = horizonclient.DefaultPublicNetClient
			networkPassphrase = network.PublicNetworkPassphrase
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmdPublicNet.AddCommand(
		createCmdSign(),
	)

	cmd.AddCommand(
		createCmdGenerateKeypair(),
		createCmdXDR(),
		&cmdTestNet,
		&cmdPublicNet,
	)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
