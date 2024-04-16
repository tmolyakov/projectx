package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyPair_Sign_VerifySuccess(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()
	msg := []byte("Hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sig.Verify(publicKey, msg))
}

func TestKeyPair_Sign_VerifyFail(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()
	msg := []byte("Hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	otherPubKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(otherPubKey, msg))
	assert.False(t, sig.Verify(publicKey, []byte("xxxxxx")))
}
