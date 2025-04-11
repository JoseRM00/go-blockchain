package blockchain

import (
	//"bytes"
	//"crypto/sha256"
	//"encoding/gob"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

// Block represents each block in the blockchain
type Block struct {
	Timestamp     time.Time
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Validator     []byte //validator's public key
	Nonce         int
}

// NewBlock: Create and return a new block by hashing the previous block and the current transactions
func NewBlock(transactions []*Transaction, prevBlockHash []byte, validator []byte) *Block {
	block := &Block{
		Timestamp:     time.Now(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Validator:     validator,
	}
	block.Hash = block.calculateHash()
	return block
}

// calculateHash() generates the hash of the block
func (b *Block) calculateHash() []byte {
	var txHashes []byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.hashTransaction()...)
	}
	hash := sha256.Sum256(bytes.Join([][]byte{
		b.PrevBlockHash,
		txHashes,
		[]byte(b.Timestamp.String()),
	}, []byte{}))

	return hash[:]
}

// serialize() converts a block into a byte slice for storage
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	//if there is an error
	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock() converts byte slice back into block
func DeserializeBlock(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	//if there is an error
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}

	return &block
}
