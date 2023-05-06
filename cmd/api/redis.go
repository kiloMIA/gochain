package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func NewRedisClient(addr, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, fmt.Errorf("failed to ping Redis server: %w", err)
	}

	return client, nil
}
