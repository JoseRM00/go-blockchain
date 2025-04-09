package storage

import (
	"log"

	"github.com/dgraph-io/badger"
)

/*
1. OpenDB will open a BadgerDB at the specified path
2. CloseDB will close the BadgerDB instance safely
3. SaveBlock will save a serialized block to the badgerDB
4. GetBlock will get a block from the badgerDB instance by its hash and deserialize it
*/

// BlockchainDB manages the blockchain storage
type BlockchainDB struct {
	DB *badger.DB
}

// OpenDB will open a BadgerDB at the specified path
func OpenDB(path string) *BlockchainDB {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		log.Panic(err)
	}
	return &BlockchainDB{DB: db}
}

// SaveBlock will save a serialized block to the badgerDB
func (bdb *BlockchainDB) SaveBlock(block *Block) error {
	return bdb.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set(block.Hash, block.Serialize())
		if err != nil {
			return err
		}
		return nil
	})
}
