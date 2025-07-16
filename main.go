package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/m-muneebrehman/Priv-BlockChain/blockchain"
)

type CommandLine struct {
	Blockchain *blockchain.Blockchain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("usage:")
	fmt.Println(" add -block Block_Data - Add a block to the blockchain")
	fmt.Println(" print -blocks - Print all blocks in the blockchain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.Blockchain.AddBlock(data)
	fmt.Println("Block added successfully!")
}

func (cli *CommandLine) printBlocks() {
	iter := cli.Blockchain.Iterator()
	for {
		block := iter.Next()
		if block == nil {
			break
		}
		fmt.Printf("Block Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s, PrevHash: %x\n", block.Data, block.PrevHash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("pow Target: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("-----------------------------------")
		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printBlocks()
	}

}

func main() {
	defer os.Exit(0)
	chain := blockchain.InitializeBlockchain()
	defer chain.Database.Close()

	cli := CommandLine{chain}
	cli.Run()

}
