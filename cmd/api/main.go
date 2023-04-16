package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	Data      string
	Hash      string
	PrevHash  string
	Nonce     int
}

const Difficulty = 3

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp.String() + block.Data + block.PrevHash + string(block.Nonce)
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

	for {
		newBlock.Nonce = rand.Intn(100000000)
		newBlock.Hash = calculateHash(newBlock)
		if isValidHash(newBlock.Hash) {
			break
		}
	}

	return newBlock
}

func isValidHash(hash string) bool {
	for i := 0; i < Difficulty; i++ {
		if hash[i] != '0' {
			return false
		}
	}
	return true
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
	if !isValidHash(newBlock.Hash) {
		return false
	}
	return true
}

func main() {
	genesisBlock := Block{0, time.Now().UTC(), "Genesis Block", "", "", 0}
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
