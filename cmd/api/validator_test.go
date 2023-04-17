package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestValidateBlockHash(t *testing.T) {
	tests := []struct {
		block     *Block
		wantValid bool
	}{
		{
			block: &Block{
				Timestamp: time.Now(),
				Pubkey:    "abc",
				Hash:      "",
			},
			wantValid: false,
		},
		{
			block: &Block{
				Timestamp: time.Now(),
				Pubkey:    "abc",
				Hash:      "invalid",
			},
			wantValid: false,
		},
		{
			block: &Block{
				Timestamp: time.Now(),
				Pubkey:    "asdsgtsey",
			},
			wantValid: true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("block=%v", test.block), func(t *testing.T) {
			got := ValidateBlockHash(test.block)
			if !reflect.DeepEqual(got, test.wantValid) {
				t.Errorf("ValidateBlockHash() = %v, want %v", got, test.wantValid)
			}
		})
	}
}
