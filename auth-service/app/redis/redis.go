package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCacheImpl struct {
	client *redis.ClusterClient
}

func New(password string, addrs ...string) *RedisCacheImpl {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	})
	return &RedisCacheImpl{
		client: rdb,
	}
}

func (rdb *RedisCacheImpl) Set(ctx context.Context, key string, value string, expirationTime time.Duration) error {
	return rdb.client.Set(ctx, key, value, expirationTime).Err()
}

func (rdb *RedisCacheImpl) Get(ctx context.Context, key string) (*string, error) {
	res, err := rdb.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &res, nil
}
