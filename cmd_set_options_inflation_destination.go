package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/txnbuild"
	"gopkg.in/AlecAivazis/survey.v1"
)

func createCmdSetOptionsInflationDestination() *cobra.Command {
	return &cobra.Command{
		Use:   "inflationdestination",
		Short: "Set the inflation destination of an account.",
		Run:   cmdSetOptionsInflationDestination,
	}
}

func cmdSetOptionsInflationDestination(cmd *cobra.Command, args []string) {
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

	var inflationAccountPublic string
	if err := survey.AskOne(&survey.Input{Message: "Inflation Account"}, &inflationAccountPublic, validatePublicKey); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Source Account:", sourceAccountPublic)
	for _, b := range sourceAccount.Balances {
		fmt.Fprintln(os.Stderr, "  Balance:", b.Balance, b.Asset.Type, b.Asset.Code, b.Asset.Issuer)
	}
	fmt.Println("Inflation Account:", inflationAccountPublic)
	inflationAccount, err := client.AccountDetail(horizonclient.AccountRequest{AccountID: inflationAccountPublic})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, b := range inflationAccount.Balances {
		fmt.Fprintln(os.Stderr, "  Balance:", b.Balance, b.Asset.Type, b.Asset.Code, b.Asset.Issuer)
	}

	tx := txnbuild.Transaction{
		SourceAccount: &sourceAccount,
		Operations: []txnbuild.Operation{
			&txnbuild.SetOptions{
				InflationDestination: txnbuild.NewInflationDestination(inflationAccountPublic),
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
