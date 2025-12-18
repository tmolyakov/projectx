package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tmolyakov/projectx/crypto"
	"github.com/tmolyakov/projectx/types"
)

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	block := randomBlock(0, types.Hash{})

	assert.Nil(t, block.Sign(privKey))
	assert.NotNil(t, block.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	block := randomBlock(0, types.Hash{})

	assert.Nil(t, block.Sign(privKey))
	assert.Nil(t, block.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	block.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, block.Verify())
}

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     uint64(time.Now().UnixNano()),
	}

	return NewBlock(header, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	block := randomBlock(height, prevBlockHash)
	privKey := crypto.GeneratePrivateKey()
	tx := randomTxWithSignature(t)
	block.AddTransaction(tx)
	assert.Nil(t, block.Sign(privKey))

	return block
}
