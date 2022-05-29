package redis

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"authentication/core/cache"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

var (
	client cache.CacheClient
)

var (
	key = "key"
	val = "val"
)

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client = New("", mr.Addr())
	code := m.Run()
	os.Exit(code)
	mr.Close()
}

func TestSet(t *testing.T) {
	err := client.Set(context.Background(), "key", "val", time.Second)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	err := client.Set(context.Background(), "key", "val", time.Second*5)
	assert.NoError(t, err)
	val, err := client.Get(context.Background(), "key")
	assert.NoError(t, err)
	assert.NotNil(t, val)
	unexistent, err := client.Get(context.Background(), "random")
	assert.NoError(t, err)
	assert.Nil(t, unexistent)
}

func TestErrorGet(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now())
	cancel()
	_, err := client.Get(ctx, "random")
	assert.Error(t, err)
}
