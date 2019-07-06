package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/xdr"
)

func createCmdSign() *cobra.Command {
	return &cobra.Command{
		Use:   "sign",
		Short: "Sign an XDR message with a secret key.",
		Run: cmdSign,
	}
}

func cmdSign(cmd *cobra.Command, args []string) {
	var inputXDR string
	if err := survey.AskOne(&survey.Input{Message: "Transaction XDR"}, &inputXDR, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var tx xdr.TransactionEnvelope
	if err := xdr.SafeUnmarshalBase64(inputXDR, &tx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if txReEncoded, err := xdr.MarshalBase64(tx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	} else if txReEncoded != inputXDR {
		fmt.Fprintln(os.Stderr, "XDR could not be fully decoded and re-encoded without losing information")
		return
	}
	Dump(os.Stderr, tx)

	hash, err := network.HashTransaction(&tx.Tx, networkPassphrase)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Fprintln(os.Stderr, "Transaction Hash:", hex.EncodeToString(hash[:]))

	var secretKey string
	if err := survey.AskOne(&survey.Input{Message: "Secret Key"}, &secretKey, validateSecretKey); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	pair := keypair.MustParse(secretKey).(*keypair.Full)

	fmt.Fprintln(os.Stderr, "Signing...")

	sig, err := pair.SignDecorated(hash[:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	tx.Signatures = append(tx.Signatures, sig)

	txBase64, err := xdr.MarshalBase64(tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Fprintln(os.Stderr, "Transaction XDR:")
	fmt.Println(txBase64)
}
