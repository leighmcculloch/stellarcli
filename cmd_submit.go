package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
	"github.com/stellar/go/xdr"
	"github.com/stellar/go/network"
)

func createCmdSubmit() *cobra.Command {
	return &cobra.Command{
		Use:   "submit",
		Short: "Submit an XDR message to the network.",
		Run: cmdSubmit,
	}
}

func cmdSubmit(cmd *cobra.Command, args []string) {
	var inputXDR string
	if err := survey.AskOne(&survey.Input{Message: "Transaction XDR"}, &inputXDR, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var tx xdr.TransactionEnvelope
	if err := xdr.SafeUnmarshalBase64(string(inputXDR), &tx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	spew.Fdump(os.Stderr, &tx)

	hash, err := network.HashTransaction(&tx.Tx, networkPassphrase)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Fprintln(os.Stderr, "Transaction Hash:", hex.EncodeToString(hash[:]))

	fmt.Fprintln(os.Stderr, "Submitting...")

	resp, err := client.SubmitTransactionXDR(inputXDR)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Fprintln(os.Stderr, "Transaction Hash:", resp.Hash)
	fmt.Fprintln(os.Stderr, "Transaction Ledger:", resp.Ledger)
}
