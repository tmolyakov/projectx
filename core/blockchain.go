package core

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

type Blockchain struct {
	storage   Storage
	lock      sync.RWMutex
	headers   []*Header
	validator Validator
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		storage: NewMemorystore(),
	}
	bc.validator = NewBlockValidator(bc)
	err := bc.AddBlockWithoutValidation(genesis)
	return bc, err
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(block *Block) error {
	if err := bc.validator.ValidateBlock(block); err != nil {
		return err
	}

	return bc.AddBlockWithoutValidation(block)
}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height >= uint32(len(bc.headers)) {
		return nil, fmt.Errorf("given header height %d exceeds blockchain height %d", height, bc.Height())
	}

	bc.lock.Lock()
	defer bc.lock.Unlock()

	return bc.headers[height], nil
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) Height() uint32 {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) AddBlockWithoutValidation(block *Block) error {
	bc.lock.Lock()
	bc.headers = append(bc.headers, block.Header)
	bc.lock.Unlock()

	logrus.WithFields(logrus.Fields{
		"height": block.Height,
		"hash":   block.Hash(BlockHasher{}),
	}).Info("adding new block")

	return bc.storage.PutBlock(block)
}
