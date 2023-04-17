package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestConnectToMongoDB(t *testing.T) {
	tests := []struct {
		uri    string
		dbName string
		err    error
	}{
		{"mongodb://localhost:27017", "testdb", nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%v", tt.uri, tt.dbName), func(t *testing.T) {
			client, err := ConnectToMongoDB(tt.uri, tt.dbName)

			if err != nil && tt.err == nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if err == nil && tt.err != nil {
				t.Errorf("Expected error: %v, but got nil", tt.err)
			}

			if client != nil {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				if err := client.Disconnect(ctx); err != nil {
					t.Errorf("Failed to disconnect from MongoDB: %v", err)
				}
			}
		})
	}
}
