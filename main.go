package main

import (
	"os"

	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
