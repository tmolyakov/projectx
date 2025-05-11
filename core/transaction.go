package core

import "github.com/tmolyakov/projectx/crypto"

type Transaction struct {
	Data []byte

	PublickKey crypto.PublicKey
	Signature  crypto.Signature
}
