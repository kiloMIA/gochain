package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func ValidateBlockHash(block *Block) bool {
	expectedHash := sha256.Sum256([]byte(fmt.Sprintf("%v%v", block.Timestamp, block.Pubkey)))
	expectedHashStr := hex.EncodeToString(expectedHash[:])
	return block.Hash == expectedHashStr
}
