package redis

import (
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

func NewClient(redisAddr string) (*redis.Client, error) {
	rDb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	if err := redisotel.InstrumentTracing(rDb); err != nil {
		return nil, err
	}

	return rDb, nil
}
