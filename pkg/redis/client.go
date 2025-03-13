package redis

import (
	"context"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

func NewClient(ctx context.Context, redisAddr string) (*redis.Client, error) {
	rDb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	if err := redisotel.InstrumentTracing(rDb); err != nil {
		return nil, err
	}

	cmd := rDb.Ping(ctx)
	_, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	return rDb, nil
}
