package main

import (
	"fmt"
	bc "github.com/amazingtensor/gogoduck/blockchain"
	"log"
	"net/http"
)

func simulate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"----Transactions to be run----\n")
	history := []string{"Alice sent Bob 10.00","Bob received 2.18 from Charlie","Charlie sent Alice 13.00"}
	for index, value := range history {
		fmt.Fprintf(w,"index:%d, transaction: %s\n", index, value)
	}
	fmt.Fprintf(w,"----Start of simulation----\n")
	egg := "gogoduck" //arbitrary string used for initialization byte vector
	initialVector := []byte(egg)
	initialTransaction := "This is a dummy genesis transaction: Alpha created 10^100^100 units!"
	blk0 := bc.NewGenesisBlock(initialVector, initialTransaction)
	fmt.Fprintf(w,"The genesis block hash: %x\n", blk0.BlockHash().Sum(nil))
	//fmt.Printf("%x\n", blk0.BlockHash().Sum(nil))
	fmt.Fprintf(w,"The genesis block transaction: %s\n",blk0.Transaction())
	//fmt.Printf("%s\n\n", blk0.Transaction())

	prevblk := blk0
	for index, transaction := range history {
		nextblk := bc.NewBlock(prevblk, transaction)
		//fmt.Println("Transaction#", index, ": Block generated: ", nextblk)
		fmt.Fprintf(w,"Transaction No. %d Block generated {\n", index)
		//fmt.Printf("   Prev:%x\n   %s\n   Cur:%x\n}\n", nextblk.PreviousHash().Sum(nil),nextblk.Transaction(),nextblk.BlockHash().Sum(nil))
		fmt.Fprintf(w,"   Prev#:%x\n   %s\n   Curr#:%x\n}\n", nextblk.PreviousHash().Sum(nil),nextblk.Transaction(),nextblk.BlockHash().Sum(nil))
		prevblk = nextblk
	}

}

func main() {
	http.HandleFunc("/", simulate)
	fmt.Println("Running server on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000",nil))
}