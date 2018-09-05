package blockchain

import (
	"github.com/ocalvet/blockchain_concepts/block"
	"github.com/ocalvet/blockchain_concepts/database"
)

type Blockchain struct {
	db database.Database
}

type Chain interface {
	Get() []block.Block
	Save([]block.Block)
	Replace([]block.Block)
}

func New(db database.Database) Blockchain {
	return Blockchain{db}
}

func (chain Blockchain) Get() []block.Block {
	blocks := []block.Block{}
	chain.db.Read("chain", "blocks", &blocks)
	return blocks
}

func (chain Blockchain) Save(blocks []block.Block) {
	chain.db.Write("chain", "blocks", blocks)
}

func (chain Blockchain) Replace(newBlocks []block.Block) {
	if len(newBlocks) > len(chain.Get()) {
		chain.Save(newBlocks)
	}
}
