package main

import (
	"encoding/hex"

	"github.com/stellar/go/txnbuild"
)

func build(tx *txnbuild.Transaction) (hash string, xdr string, err error) {
	if err := tx.Build(); err != nil {
		return "", "", err
	}

	if err := tx.Sign(); err != nil {
		return "", "", err
	}

	xdr, err = tx.Base64()
	if err != nil {
		return "", "", err
	}

	hashBytes, err := tx.Hash()
	if err != nil {
		return "", "", err
	}
	hash = hex.EncodeToString(hashBytes[:])

	return hash, xdr, nil
}
