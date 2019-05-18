package main

import (
	"fmt"

	"github.com/stellar/go/amount"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/xdr"
)

func validateSecretKey(v interface{}) error {
	pair, err := keypair.Parse(v.(string))
	if err != nil {
		return fmt.Errorf("input is not a secret key: %v", err)
	}
	if _, ok := pair.(*keypair.Full); !ok {
		return fmt.Errorf("input is a public key, secret key required")
	}
	return nil
}

func validatePublicKey(v interface{}) error {
	pair, err := keypair.Parse(v.(string))
	if err != nil {
		return fmt.Errorf("input is not a public key: %v", err)
	}
	if _, ok := pair.(*keypair.FromAddress); !ok {
		return fmt.Errorf("input is a secret key, public key required")
	}
	return nil
}

func validateAmount(v interface{}) error {
	_, err := amount.Parse(v.(string))
	if err != nil {
		return fmt.Errorf("invalid amount: %v", err)
	}
	return nil
}

func validateTransactionEnvelopeXDR(v interface{}) error {
	var tx xdr.TransactionEnvelope
	return xdr.SafeUnmarshalBase64(v.(string), &tx)
}
