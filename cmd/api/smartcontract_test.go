package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestNewBlock(t *testing.T) {

	testCases := []struct {
		pubkey       string
		expectedHash string
	}{
		{"publickey1", "hash value 1"},
		{"publickey2", "hash value 2"},
		{"publickey3", "hash value 3"},
	}

	for _, testCase := range testCases {
		block := NewBlock(testCase.pubkey)

		hash := sha256.Sum256([]byte(fmt.Sprintf("%v%v", block.Timestamp, block.Pubkey)))
		expectedHash := hex.EncodeToString(hash[:])

		if block.Hash != expectedHash {
			t.Errorf("NewBlock(%v) = %v; expected %v", testCase.pubkey, block.Hash, expectedHash)
		}
	}
}
