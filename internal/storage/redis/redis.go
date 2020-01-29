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

const apqPrefix = "apq:"

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

func (c *Cache) Add(ctx context.Context, hash string, query string) {
	c.client.Set(apqPrefix+hash, query, c.ttl)
}

func (c *Cache) Get(ctx context.Context, hash string) (string, bool) {
	s, err := c.client.Get(apqPrefix + hash).Result()
	if err != nil {
		return "did not find APQ in cache", false
	}
	return s, true
}
