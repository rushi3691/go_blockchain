package main

import (
	"fmt"

	"github.com/rushi3691/go_blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")
	for _, block := range chain.Blocks {
		pow := blockchain.NewProof(block)
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		println("PoW: ", pow.Validate())

		println()
	}
}
