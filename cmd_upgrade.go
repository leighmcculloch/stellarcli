package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/rubblelabs/ripple/crypto"
	"gopkg.in/AlecAivazis/survey.v1"
	"github.com/stellar/go/keypair"
)

func createCmdUpgrade() *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade",
		Short: "Convert old Stellar secret key to new stellar secret key for use with new Stellar network.",
		Run: cmdUpgrade,
	}
}

func cmdUpgrade(cmd *cobra.Command, args []string) {
	var oldSecretKey string
	if err := survey.AskOne(&survey.Input{Message: "Old Stellar Network Secret Key"}, &oldSecretKey, nil); err != nil {
		fmt.Println(err)
		return
	}
	const oldAlphabet = "gsphnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCr65jkm8oFqi1tuvAxyz"
	oldSecretKeyRaw, err := crypto.Base58Decode(oldSecretKey, oldAlphabet)
	if err != nil {
		fmt.Println(err)
		return
	}

	var rawSeed [32]byte
	copy(rawSeed[:], oldSecretKeyRaw[1:len(oldSecretKeyRaw)-2])

	kp, err := keypair.FromRawSeed(rawSeed)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("New Public:", kp.Address())
	fmt.Println("New Secret:", kp.Seed())
}
