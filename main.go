package main

import (
	"github.com/lol-chain/kek-faucet/cmd"
)

//go:generate npm run build
func main() {
	cmd.Execute()
}
