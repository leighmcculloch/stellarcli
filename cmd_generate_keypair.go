package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stellar/go/keypair"
)

func createCmdGenerateKeypair() *cobra.Command {
	return &cobra.Command{
		Use:   "generatekeypair",
		Short: "Generate keypair that can be used as an account or keys for signing.",
		Run:   cmdGenerateKeypair,
	}
}

func cmdGenerateKeypair(cmd *cobra.Command, args []string) {
	pair, err := keypair.Random()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Public:", pair.Address())
	fmt.Println("Secret:", pair.Seed())
}
