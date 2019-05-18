package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/txnbuild"
	"gopkg.in/AlecAivazis/survey.v1"
)

func createCmdMergeAccount() *cobra.Command {
	return &cobra.Command{
		Use:   "mergeaccount",
		Short: "Merge account into another account.",
		Run:   cmdMergeAccountRun,
	}
}

func cmdMergeAccountRun(cmd *cobra.Command, args []string) {
	var sourceAccountPublic string
	if err := survey.AskOne(&survey.Input{Message: "Source Account"}, &sourceAccountPublic, validatePublicKey); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	sourceAccount, err := client.AccountDetail(horizonclient.AccountRequest{AccountID: sourceAccountPublic})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var destinationAccountPublic string
	if err := survey.AskOne(&survey.Input{Message: "Destination Account"}, &destinationAccountPublic, validatePublicKey); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Source Account:", sourceAccountPublic)
	for _, b := range sourceAccount.Balances {
		fmt.Fprintln(os.Stderr, "  Balance:", b.Balance, b.Asset.Type, b.Asset.Code, b.Asset.Issuer)
	}
	fmt.Println("Destination Account:", destinationAccountPublic)
	destinationAccount, err := client.AccountDetail(horizonclient.AccountRequest{AccountID: destinationAccountPublic})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, b := range destinationAccount.Balances {
		fmt.Fprintln(os.Stderr, "  Balance:", b.Balance, b.Asset.Type, b.Asset.Code, b.Asset.Issuer)
	}

	tx := txnbuild.Transaction{
		SourceAccount: &sourceAccount,
		Operations: []txnbuild.Operation{
			&txnbuild.AccountMerge{
				Destination:   destinationAccountPublic,
				SourceAccount: &sourceAccount,
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
