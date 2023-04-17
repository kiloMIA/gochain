package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func ConnectToRedis(addr string, password string, db int) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis:", pong)
	return client, nil
}
