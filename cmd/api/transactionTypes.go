package main

import "fmt"

func handleTypeAA01(tx Transaction) Transaction {
	txtype := tx.HashValue
	fmt.Println(txtype)
	fmt.Println("Handling type AA01 transaction:", tx)
	return tx
}

func handleTypeAA04(tx Transaction) Transaction {
	txtype := tx.HashValue
	fmt.Println(txtype)
	fmt.Println("Handling type AA04 transaction:", tx)
	return tx
}

func handleTypeAA05(tx Transaction) Transaction {
	txtype := tx.HashValue
	fmt.Println(txtype)
	fmt.Println("Handling type AA05 transaction:", tx)
	return tx
}
func handleUnknownType(tx Transaction) Transaction {
	txtype := tx.HashValue
	fmt.Println(txtype)
	fmt.Println("Unknown transaction type:", tx.HashValue)
	return tx
}
