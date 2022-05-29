package cache

import (
	"context"
	"time"
)

type CacheClient interface {
	Set(ctx context.Context, key string, value string, expirationTime time.Duration) error
	Get(ctx context.Context, key string) (*string, error)
}
