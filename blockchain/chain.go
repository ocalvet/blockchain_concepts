package blockchain

import (
	"github.com/ocalvet/blockchain_concepts/block"
)

type Blockchain struct {
}

type Chain interface {
	Get() []block.Block
	Save([]block.Block)
	Replace([]block.Block)
}

func New(location string) Blockchain {
	return Blockchain{}
}

func (chain Blockchain) Get() []block.Block {
	return []block.Block{}
}

func (chain Blockchain) Save([]block.Block) {
}

func (chain Blockchain) Replace(newBlocks []block.Block) {
	if len(newBlocks) > len(chain.Get()) {
		// TODO - save longest chain
		chain.Save(newBlocks)
	}
}
