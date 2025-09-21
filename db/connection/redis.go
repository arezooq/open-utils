package connection

import (
	"context"

	"github.com/arezooq/open-utils/errors"
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
    Addr     string
    Password string
    DB       int
    PoolSize int
}

func ConnectRedis(ctx context.Context, cfg RedisConfig) (*redis.Client, error) {
    client := redis.NewClient(&redis.Options{
        Addr:     cfg.Addr,
        Password: cfg.Password,
        DB:       cfg.DB,
        PoolSize: cfg.PoolSize,
    })

    if err := client.Ping(ctx).Err(); err != nil {
        return nil, errors.New("failed to connect to redis: %w", err.Error(), 500)
    }

    return client, nil
}