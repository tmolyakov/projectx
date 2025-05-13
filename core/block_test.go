package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tmolyakov/projectx/crypto"
	"github.com/tmolyakov/projectx/types"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     uint64(time.Now().UnixNano()),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func randomBlockWithSignature(t *testing.T, height uint32) *Block {
	block := randomBlock(height)
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, block.Sign(privKey))

	return block
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	block := randomBlock(0)

	assert.Nil(t, block.Sign(privKey))
	assert.NotNil(t, block.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	block := randomBlock(0)

	assert.Nil(t, block.Sign(privKey))
	assert.Nil(t, block.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	block.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, block.Verify())
}
