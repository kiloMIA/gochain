package main

import (
	"fmt"
	"testing"
)

func TestConnectToRedis(t *testing.T) {
	tests := []struct {
		addr     string
		password string
		db       int
		wantErr  bool
	}{
		{
			addr:     "localhost:6379",
			password: "",
			db:       0,
			wantErr:  false,
		},
		{
			addr:     "localhost:9999",
			password: "",
			db:       0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("ConnectToRedis(%v, %v, %v)", tt.addr, tt.password, tt.db), func(t *testing.T) {
			got, err := ConnectToRedis(tt.addr, tt.password, tt.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectToRedis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				defer got.Close()
				gotPong, err := got.Ping().Result()
				if err != nil {
					t.Errorf("ConnectToRedis() error when pinging Redis = %v", err)
					return
				}
				if gotPong != "PONG" {
					t.Errorf("ConnectToRedis() got pong = %v, want PONG", gotPong)
					return
				}
			}
		})
	}
}
