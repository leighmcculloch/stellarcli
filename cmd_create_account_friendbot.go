package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stellar/go/clients/horizonclient"
	"gopkg.in/AlecAivazis/survey.v1"
)

func createCmdCreateAccountFriendbot() *cobra.Command {
	return &cobra.Command{
		Use:   "createaccountfriendbot",
		Short: "Create account with the testnet's Friendbot.",
		Run:   cmdCreateAccountFriendbotRun,
	}
}

func cmdCreateAccountFriendbotRun(cmd *cobra.Command, args []string) {
	var destinationAccountPublic string
	if err := survey.AskOne(&survey.Input{Message: "Destination Account"}, &destinationAccountPublic, validatePublicKey); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := client.Fund(destinationAccountPublic); err != nil {
		fmt.Println(err)
		return
	}

	account, err := client.AccountDetail(horizonclient.AccountRequest{AccountID: destinationAccountPublic})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, b := range account.Balances {
		fmt.Println("Balance:", b.Balance, b.Asset.Type, b.Asset.Code, b.Asset.Issuer)
	}
	for _, s := range account.Signers {
		fmt.Println("Signer:", "Weight:", s.Weight, "Key:", s.Key)
	}
}
