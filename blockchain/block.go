package blockchain

/*
block data structure has a hash based on previous hash and current transaction
 */
type Block struct {
	previousHash int
	transaction string
	blockHash int
}

func (b *Block) BlockHash() int {
	return b.blockHash
}

func (b *Block) PreviousHash() int {
	return b.previousHash
}

func (b *Block) Transaction() string {
	return b.transaction
}

func NewGenesisBlock(initialVector int, initialTransaction string) *Block{
	blockHash := initialVector + len(initialTransaction)
	return &Block{initialVector, initialTransaction, blockHash}
}

func NewBlock(previousBlock Block, currentTransaction string) *Block {
	blockHash := previousBlock.BlockHash() + len(currentTransaction)
	return &Block{previousBlock.BlockHash(), currentTransaction, blockHash}
}



