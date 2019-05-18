package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/txnbuild"
	"gopkg.in/AlecAivazis/survey.v1"
)

func createCmdCreateAccount() *cobra.Command {
	return &cobra.Command{
		Use:   "createaccount",
		Short: "Create account.",
		Run:   cmdCreateAccountRun,
	}
}

func cmdCreateAccountRun(cmd *cobra.Command, args []string) {
	var sourceAccountPublic string
	if err := survey.AskOne(&survey.Input{Message: "Source Account"}, &sourceAccountPublic, validatePublicKey); err != nil {
		fmt.Println(err)
		return
	}
	sourceAccount, err := client.AccountDetail(horizonclient.AccountRequest{AccountID: sourceAccountPublic})
	if err != nil {
		fmt.Println(err)
		return
	}

	var destinationAccountPublic string
	if err := survey.AskOne(&survey.Input{Message: "Destination Account"}, &destinationAccountPublic, validatePublicKey); err != nil {
		fmt.Println(err)
		return
	}

	var amount string
	if err := survey.AskOne(&survey.Input{Message: "Destination Account Starting Amount"}, &amount, validateAmount); err != nil {
		fmt.Println(err)
		return
	}

	tx := txnbuild.Transaction{
		SourceAccount: &sourceAccount,
		Operations: []txnbuild.Operation{
			&txnbuild.CreateAccount{
				Destination: destinationAccountPublic,
				Amount:      amount,
			},
		},
		Timebounds: txnbuild.NewTimeout(300),
		Network:    networkPassphrase,
	}

	hash, xdr, err := build(&tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Fprintln(os.Stderr, "Transaction Hash:", hash)
	fmt.Fprintln(os.Stderr, "Transaction XDR:")
	fmt.Println(xdr)
}
