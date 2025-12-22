package network

import (
	"github.com/tmolyakov/projectx/core"
	"github.com/tmolyakov/projectx/types"
)

type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}



func p *TxPool) Has(hash types.Hash) bool {
	_, exists := p.transactions[hash]
	return exists
}

func (p *TxPool) Len() int {
	return len(p.transactions)
}

func (p *TxPool) Flush() {
	p.transactions = make(map[types.Hash]*core.Transaction)
}
