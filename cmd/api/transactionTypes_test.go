package main

import (
	"reflect"
	"testing"
)

func TestHandleTypeAA01(t *testing.T) {
	testCases := []struct {
		name     string
		tx       Transaction
		expected Transaction
	}{
		{
			name: "Valid transaction",
			tx: Transaction{
				Pubkey:    "1234567890abcdef",
				HashValue: "AA01abcdef123456",
				Detail: TypeAA01{
					Foo: "foo",
					Bar: "bar",
				},
			},
			expected: Transaction{
				Pubkey:    "1234567890abcdef",
				HashValue: "AA01abcdef123456",
				Detail: TypeAA01{
					Foo: "foo",
					Bar: "bar",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := handleTypeAA01(tc.tx)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
func TestHandleTypeAA04(t *testing.T) {
	testCases := []struct {
		name     string
		tx       Transaction
		expected Transaction
	}{
		{
			name: "Valid transaction",
			tx: Transaction{
				Pubkey:    "1234567890abcdef",
				HashValue: "AA04abcdef123456",
				Detail: TypeAA04{
					Baz: "baz",
				},
			},
			expected: Transaction{
				Pubkey:    "1234567890abcdef",
				HashValue: "AA04abcdef123456",
				Detail: TypeAA04{
					Baz: "baz",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := handleTypeAA04(tc.tx)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
func TestHandleTypeAA05(t *testing.T) {
	testCases := []struct {
		name     string
		tx       Transaction
		expected Transaction
	}{
		{
			name: "Valid transaction",
			tx: Transaction{
				Pubkey:    "1234567890abcdef",
				HashValue: "AA05abcdef123456",
				Detail: TypeAA05{
					Qux: "qux",
				},
			},
			expected: Transaction{
				Pubkey:    "1234567890abcdef",
				HashValue: "AA05abcdef123456",
				Detail: TypeAA05{
					Qux: "qux",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := handleTypeAA05(tc.tx)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
