package main

import (
	"os"

	"github.com/m-muneebrehman/Priv-BlockChain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
