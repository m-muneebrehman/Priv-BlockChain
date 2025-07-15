package main

import (
	"fmt"
	"strconv"

	"github.com/m-muneebrehman/Priv-BlockChain/blockchain"
)

func main() {
	chain := blockchain.InitializeBlockchain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	for _, block := range chain.Blocks {
		fmt.Printf("Block Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s, PrevHash: %x\n", block.Data, block.PrevHash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("pow Target: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("-----------------------------------")
	}

}
