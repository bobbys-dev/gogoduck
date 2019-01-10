package main

import (
	"fmt"
	bc "github.com/amazingtensor/gogoduck/blockchain"
)


func main() {
	fmt.Println("----Transactions to be run----")

	history := []string{"Alice sent Bob 10.00","Bob received 2.15 from Charlie","Charlie sent Alice 13.00"}
	for index, value := range history {
		fmt.Println("index:", index, "value:", value)
	}
	fmt.Println("----Start of simulation----")
	initialVector := 0
	initialTransaction := "This is a placeholder for the imaginary genesis transaction."
	blk0 := bc.NewGenesisBlock(initialVector, initialTransaction)
	fmt.Println("The genesis block is:", *blk0)


	prevblk := *blk0
	for index, transaction := range history {
		nextblk := bc.NewBlock(prevblk, transaction)
		fmt.Println("Transaction#", index, ": Block generated: ", *nextblk)
		prevblk = *nextblk
	}


}