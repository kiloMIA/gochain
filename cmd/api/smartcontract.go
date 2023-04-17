package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Timestamp time.Time
	Hash      string
	Pubkey    string
}

func NewBlock(pubkey string) *Block {
	block := &Block{
		Timestamp: time.Now(),
		Pubkey:    pubkey,
	}

	hash := sha256.Sum256([]byte(fmt.Sprintf("%v%v", block.Timestamp, block.Pubkey)))
	block.Hash = hex.EncodeToString(hash[:])

	return block
}
