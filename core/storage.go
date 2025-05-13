package core

type Storage interface {
	PutBlock(block *Block) error
}

type MemoryStore struct {
}

func NewMemorystore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) PutBlock(b *Block) error {
	return nil
}
