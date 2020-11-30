package main

import (
	"os"

	"github.com/bytejedi/ledger/cmd/ledgerd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
