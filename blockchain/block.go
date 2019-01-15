package blockchain

import (
	"crypto/sha256"
	"hash"
)

/*
block data structure has a hash based on previous hash and current transaction
 */
type Block struct {
	previousHash hash.Hash
	transaction string
	blockHash hash.Hash
}

func ( b* Block) PreviousHash() hash.Hash {
	return b.previousHash
}

func (b* Block) Transaction() string {
	return b.transaction
}

func (b* Block) BlockHash() hash.Hash {
	return b.blockHash
}

func NewGenesisBlock(initialVector []byte, initialTransaction string) Block{
	var hashStream []byte

	//copy over initial vector'shash
	initialHash := sha256.New()
	initialHash.Write(initialVector)
	for _, B := range initialHash.Sum(nil) {
		hashStream = append(hashStream,B)
	}

	//copy over hash of initial transaction
	cur := sha256.Sum256([]byte(initialTransaction))
	for _, B := range cur {
		hashStream = append(hashStream,B)
	}

	//digest hash
	blockHash := sha256.New()
	blockHash.Write(hashStream)

	return Block{initialHash, initialTransaction, blockHash}
}

func NewBlock(previousBlock Block, currentTransaction string) Block {
	var hashStream []byte

	//copy over previous hash
	for _, B := range previousBlock.BlockHash().Sum(nil) {
		hashStream = append(hashStream,B)
	}

	//copy over hash of current transaction
	cur := sha256.Sum256([]byte(currentTransaction))
	for _, B := range cur {
		hashStream = append(hashStream,B)
	}

	//digest current block's hash
	blockHash := sha256.New()
	blockHash.Write(hashStream)
	return Block{previousBlock.BlockHash(), currentTransaction, blockHash}
}