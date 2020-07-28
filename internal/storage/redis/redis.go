package redis

import (
	"context"
	"fmt"
	"time"

	r "github.com/go-redis/redis/v7"
)

type Cache struct {
	client r.UniversalClient
	ttl    time.Duration
}

func NewCache(redisAddress string, ttl time.Duration) (*Cache, error) {
	client := r.NewClient(&r.Options{
		Addr: redisAddress,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("error pinging redis %s", err)
	}

	return &Cache{client: client, ttl: ttl}, nil
}

func (c *Cache) Get(ctx context.Context, key string) (string, bool) {
	val, err := c.client.Get(key).Result()
	if err != nil {
		return fmt.Sprintf("did not find %s in cache", key), false
	}
	return val, true
}

func (c *Cache) Set(ctx context.Context, key string, value string) {
	c.client.Set(key, value, c.ttl)
}
