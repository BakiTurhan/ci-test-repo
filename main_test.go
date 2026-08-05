package main

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestRedisConnectionInCI(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer client.Close()

	err := client.Ping(context.Background()).Err()
	if err != nil {
		t.Fatalf("Github Actions üzerinde Redis'e ulaşılamadı! Hata: %v", err)
	}

	t.Log("Github Actions üzerinde Redis başarıyla ayağa kalktı ve yanıt veriyor!")
}
