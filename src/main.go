package main

import (
	"fmt"
	"strconv"
)

func main() {

	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Abhi")
	bc.AddBlock("Send 2 BTC to Abhi")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofofwork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
