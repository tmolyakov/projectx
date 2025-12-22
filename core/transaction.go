package core

import (
	"crypto/sha256"
	"fmt"

	"github.com/tmolyakov/projectx/crypto"
	"github.com/tmolyakov/projectx/types"
)

type Transaction struct {
	Data []byte

	From      crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) Hash() types.Hash {
	return types.Hash(sha256.Sum256(tx.Data))
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.From = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("signature is nil")
	}

	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("signature verification failed")
	}

	return nil
}
