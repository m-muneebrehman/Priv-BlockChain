package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Blockchain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func GenesisBlock() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitializeBlockchain() *Blockchain {
	genesis := GenesisBlock()
	return &Blockchain{[]*Block{genesis}}
}

func main() {
	blockchain := InitializeBlockchain()
	blockchain.AddBlock("First Block after Genesis")
	blockchain.AddBlock("Second Block after Genesis")
	for _, block := range blockchain.blocks {
		fmt.Printf("Block Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s, PrevHash: %x\n", block.Data, block.PrevHash)
	}

}
