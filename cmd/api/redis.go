package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// NewRedisClient returns a new Redis client with the specified options.
func NewRedisClient(addr, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Ping the Redis server to test the connection
	if err := client.Ping().Err(); err != nil {
		return nil, fmt.Errorf("failed to ping Redis server: %w", err)
	}

	return client, nil
}
