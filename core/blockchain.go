package core

type Blockchain struct {
	storage   Storage
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

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) AddBlockWithoutValidation(block *Block) error {
	bc.headers = append(bc.headers, block.Header)
	return bc.storage.PutBlock(block)
}
