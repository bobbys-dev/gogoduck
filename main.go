package main

import (
	"fmt"
	bc "github.com/amazingtensor/gogoduck/blockchain"
)

func main() {
	fmt.Println("----Transactions to be run----")

	history := []string{"Alice sent Bob 10.00","Bob received 2.18 from Charlie","Charlie sent Alice 13.00"}
	for index, value := range history {
		fmt.Println("index:", index, "value:", value)
	}
	fmt.Println("----Start of simulation----")
	egg := "gogoduck" //arbitrary string used for initialization byte vector
	initialVector := []byte(egg)
	initialTransaction := "This is a dummy genesis transaction: Alpha created 10^100^100 units!"
	blk0 := bc.NewGenesisBlock(initialVector, initialTransaction)
	fmt.Println("The genesis block hash:")
	fmt.Printf("%x\n", blk0.BlockHash().Sum(nil))
	fmt.Println("The genesis block transaction:")
	fmt.Printf("%s\n\n", blk0.Transaction())

	prevblk := blk0
	for index, transaction := range history {
		nextblk := bc.NewBlock(prevblk, transaction)
		//fmt.Println("Transaction#", index, ": Block generated: ", nextblk)
		fmt.Println("Transaction#", index, ": Block generated: {")
		fmt.Printf("   Prev:%x\n   %s\n   Cur:%x\n}\n", nextblk.PreviousHash().Sum(nil),nextblk.Transaction(),nextblk.BlockHash().Sum(nil))
		prevblk = nextblk
	}
}