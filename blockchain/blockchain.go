package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/rushi3691/go_blockchain/utils"
)

const (
	dbPath = "./tmp/blocks"
)

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
	Blocks   []*Block
}

func InitBlockChain() *Blockchain {
	// return &Blockchain{[]*Block{Genesis()}}
	var lastHash []byte
	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath
	db, err := badger.Open(opts)
	utils.HandleErr(err)
	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			utils.HandleErr(err)
			err = txn.Set([]byte("lh"), genesis.Hash)
			lastHash = genesis.Hash
			return err
		}
		return nil
	},
	)
	utils.HandleErr(err)

}
