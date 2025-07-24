package main

import (
	"fmt"
	"os"

	"github.com/m-muneebrehman/Priv-BlockChain/blockchain"
)

func main() {
	defer os.Exit(0)

	// Test blockchain creation with a fixed address for testing
	nodeID := "3000"
	testAddress := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa" // A known valid Bitcoin address format

	fmt.Println("Creating blockchain...")
	chain := blockchain.InitBlockChain(testAddress, nodeID)
	defer chain.Database.Close()

	fmt.Println("Blockchain created successfully!")

	// Test printing the blockchain
	fmt.Println("Printing blockchain...")
	iter := chain.Iterator()
	for {
		block := iter.Next()
		if block == nil {
			break
		}
		fmt.Printf("Block Hash: %x\n", block.Hash)
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Transactions: %d\n", len(block.Transactions))
		fmt.Println("---")

		if len(block.PrevHash) == 0 {
			break
		}
	}
}
