package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []Transaction
	Hash         string
	PrevHash     string
	Nonce        int
}

const Difficulty = 3

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp.String() + fmt.Sprintf("%v", block.Transactions) + block.PrevHash + strconv.Itoa(block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock Block, transactions []Transaction) Block {
	var newBlock Block
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().UTC()
	newBlock.Transactions = transactions
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
func isTransactionValid(transaction Transaction, balanceMap map[string]float64) bool {
	if balanceMap[transaction.Sender] < transaction.Amount {
		return false
	}
	return true
}
func resolveConflicts(chain1, chain2 []Block) []Block {
	if len(chain1) > len(chain2) {
		return chain1
	} else if len(chain1) < len(chain2) {
		return chain2
	} else {

		if chain1[len(chain1)-1].Hash == chain2[len(chain2)-1].Hash {
			return chain1
		} else {

			if chain1[len(chain1)-1].Nonce > chain2[len(chain2)-1].Nonce {
				return chain1
			} else {
				return chain2
			}
		}
	}
}

func main() {
	genesisBlock := Block{0, time.Now().UTC(), []Transaction{}, "", "", 0}
	genesisBlock.Hash = calculateHash(genesisBlock)
	blockchain := []Block{genesisBlock}

	transactions1 := []Transaction{
		{"sender1", "receiver1", 10.0},
		{"sender2", "receiver2", 5.0},
	}
	block1 := generateBlock(blockchain[0], transactions1)
	if isBlockValid(block1, blockchain[0]) {
		blockchain = append(blockchain, block1)
	}

	transactions2 := []Transaction{
		{"sender1", "receiver3", 7.0},
		{"sender3", "receiver2", 3.0},
	}
	block2 := generateBlock(blockchain[1], transactions2)
	if isBlockValid(block2, blockchain[1]) {
		blockchain = append(blockchain, block2)
	}

	fmt.Println(blockchain)

	balanceMap := map[string]float64{
		"sender1": 20.0,
		"sender2": 10.0,
	}
	transaction := Transaction{"sender1", "receiver1", 5.0}
	if isTransactionValid(transaction, balanceMap) {
		fmt.Println("Transaction is valid")
	} else {
		fmt.Println("Transaction is invalid")
	}

	chain1 := []Block{genesisBlock, block1}
	chain2 := []Block{genesisBlock, block2}
	newChain := resolveConflicts(chain1, chain2)
	fmt.Println(newChain)
}
