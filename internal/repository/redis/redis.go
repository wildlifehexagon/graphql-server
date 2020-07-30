package redis

import (
	"fmt"
	"time"

	"github.com/dictyBase/graphql-server/internal/repository"
	r "github.com/go-redis/redis/v7"
)

type RedisStorage struct {
	client *r.Client
}

func NewCache(addr string) (repository.Repository, error) {
	client := r.NewClient(&r.Options{
		Addr: addr,
	})
	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("error pinging redis %s", err)
	}
	return &RedisStorage{client: client}, nil
}

func (rs *RedisStorage) Get(key string) (string, error) {
	val, err := rs.client.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("did not find %s in cache %s", key, err)
	}
	return val, nil
}

func (rs *RedisStorage) Set(key string, value string) error {
	return rs.client.Set(key, value, 24*time.Hour).Err()
}

func (rs *RedisStorage) Exists(key string) (bool, error) {
	h, err := rs.client.Exists(key).Result()
	if err != nil {
		return false, err
	}
	if h == 0 {
		return false, nil
	}
	return true, nil
}

func (rs *RedisStorage) HGet(hash, key string) (string, error) {
	val, err := rs.client.HGet(hash, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (rs *RedisStorage) HSet(hash, key, value string) error {
	return rs.client.HSet(hash, key, value).Err()
}

func (rs *RedisStorage) HExists(key, field string) (bool, error) {
	h, err := rs.client.HExists(key, field).Result()
	if err != nil {
		return false, err
	}
	return h, err
}
