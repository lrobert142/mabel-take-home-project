package main

import "fmt"

func main() {
	/*
		 TODO:
		 Read `mable_account_balances.csv` into Account structs <-- Maybe we need a special CSV reader?
		 Read `mable_transactions.csv` into Transaction structs <-- Maybe we need a special CSV reader?
		 For each transaction:
		 	Find account with ID <from> (else return err)
			Find account with ID <to> (else return err)
			Transfer amount from <from> to <to> (internal handling: if amount would exceed what's available in <from>, return err)
	*/
	fmt.Println("Hello, world.")
}
