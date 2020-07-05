package main

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/ripemd160"

	"github.com/rubblelabs/ripple/crypto"
	"github.com/spf13/cobra"
	"github.com/stellar/go/keypair"
	"gopkg.in/AlecAivazis/survey.v1"
)

func createCmdUpgrade() *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade",
		Short: "Convert old Stellar secret key to new stellar secret key for use with new Stellar network.",
		Run:   cmdUpgrade,
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

	sk := ed25519.NewKeyFromSeed(rawSeed[:])
	pk := sk.Public()
	pkBytes := []byte(pk.(ed25519.PublicKey))
	pkSHA256 := sha256.Sum256(pkBytes)
	md := ripemd160.New()
	md.Write(pkSHA256[:])
	pkRIPE160SHA256 := md.Sum(nil)
	oldAccountIDBytes := append([]byte{0}, pkRIPE160SHA256...)

	oldPublicKey := crypto.Base58Encode(oldAccountIDBytes, oldAlphabet)
	fmt.Println("Old Public:", oldPublicKey)

	kp, err := keypair.FromRawSeed(rawSeed)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("New Public:", kp.Address())
	fmt.Println("New Secret:", kp.Seed())
}
