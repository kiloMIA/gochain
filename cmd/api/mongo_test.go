package main

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestNewMongoDBClient(t *testing.T) {
	tests := []struct {
		name string
		uri  string
		want *mongo.Client
		err  bool
	}{
		{
			name: "Successful connection",
			uri:  "mongodb://localhost:27017",
			want: &mongo.Client{},
			err:  false,
		},
		{
			name: "Invalid URI",
			uri:  "invalid-uri",
			want: nil,
			err:  true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMongoDBClient(ctx, tt.uri)
			if (err != nil) != tt.err {
				t.Errorf("NewMongoDBClient() error = %v, wantErr %v", err, tt.err)
				return
			}

			if tt.want != nil && got == nil {
				t.Errorf("NewMongoDBClient() got = nil, want %v", tt.want)
				return
			}

			if tt.want != nil && got.Database("test").Name() != tt.want.Database("test").Name() {
				t.Errorf("NewMongoDBClient() got = %v, want %v", got.Database("test").Name(), tt.want.Database("test").Name())
			}
		})
	}
}
