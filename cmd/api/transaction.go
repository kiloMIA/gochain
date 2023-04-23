package main

import (
	"time"
)

type Transactions struct {
	List []Transaction `json:"transactions"`
}

type Transaction struct {
	Timestamp time.Time   `json:"timestamp"`
	Pubkey    string      `json:"pubkey"`
	HashValue string      `json:"hashvalue"`
	Detail    interface{} `json:"detail"`
}

func handleTransaction(tx Transaction) {
	txType := tx.HashValue[:4]

	switch txType {
	case "AA01":
		handleTypeAA01(tx)
	case "AA04":
		handleTypeAA04(tx)
	case "AA05":
		handleTypeAA05(tx)
	default:
		handleUnknownType(tx)
	}
}
