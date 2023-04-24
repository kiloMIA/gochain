package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	testCases := []struct {
		name         string
		requestBody  string
		expectedCode int
	}{
		{
			name:         "Valid transaction",
			requestBody:  `{"List":[{"Timestamp":"2022-04-23T10:00:00Z","Pubkey":"1234567890abcdef","HashValue":"AA01abcdef123456","Detail":{"Foo":"foo","Bar":"bar"}}]}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid JSON",
			requestBody:  `{"List":[{]}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/create_tx", bytes.NewBufferString(tc.requestBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(createTransaction)
			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectedCode {
				t.Errorf("Expected %d, but got %d", tc.expectedCode, rr.Code)
			}
		})
	}
}
