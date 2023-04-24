package main

import (
	"testing"
	"time"
)

func TestNewRedisClient(t *testing.T) {
	tests := []struct {
		name     string
		addr     string
		password string
		db       int
		wantErr  bool
	}{
		{
			name:     "Valid configuration",
			addr:     "localhost:6379",
			password: "",
			db:       0,
			wantErr:  false,
		},
		{
			name:     "Invalid password",
			addr:     "localhost:6379",
			password: "invalid-password",
			db:       0,
			wantErr:  true,
		},
		{
			name:     "Invalid address",
			addr:     "invalid-address:6379",
			password: "",
			db:       0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRedisClient(tt.addr, tt.password, tt.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRedisClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// Test using the client
	client, err := NewRedisClient("localhost:6379", "", 0)
	if err != nil {
		t.Errorf("failed to create Redis client: %s", err)
	}

	defer client.Close()

	err = client.Set("key", "value", time.Hour).Err()
	if err != nil {
		t.Errorf("failed to set key: %s", err)
	}
}
