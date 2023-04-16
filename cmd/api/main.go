package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	Data      string
	Hash      string
	PrevHash  string
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp.String() + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock Block, data string) Block {
	var newBlock Block
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().UTC()
	newBlock.Data = data
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func isBlockValid(newBlock, prevBlock Block) bool {
	if prevBlock.Index+1 != newBlock.Index {
		return false
	}
	if prevBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func main() {
	genesisBlock := Block{0, time.Now().UTC(), "Genesis Block", "", ""}
	genesisBlock.Hash = calculateHash(genesisBlock)
	blockchain := []Block{genesisBlock}

	block1 := generateBlock(blockchain[0], "Block 1")
	if isBlockValid(block1, blockchain[0]) {
		blockchain = append(blockchain, block1)
	}
	block2 := generateBlock(blockchain[1], "Block 2")
	if isBlockValid(block2, blockchain[1]) {
		blockchain = append(blockchain, block2)
	}

	fmt.Println(blockchain)
}
