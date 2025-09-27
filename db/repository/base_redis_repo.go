package repository

import (
	"context"
	"time"

	"github.com/arezooq/open-utils/errors"
	"github.com/redis/go-redis/v9"
)

type BaseRedisRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewBaseRedisRepository(client *redis.Client, ctx context.Context) *BaseRedisRepository {
	return &BaseRedisRepository{
		client: client,
		ctx:    ctx,
	}
}
func (r *BaseRedisRepository) Set(key string, value interface{}, ttl time.Duration) error {
	if err := r.client.Set(r.ctx, key, value, ttl).Err(); err != nil {
		return errors.ErrRedisOperation
	}
	return nil
}

func (r *BaseRedisRepository) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", errors.ErrNotFound
	}
	if err != nil {
		return "", errors.ErrRedisOperation
	}
	return val, nil
}

func (r *BaseRedisRepository) Delete(key string) error {
	if err := r.client.Del(r.ctx, key).Err(); err != nil {
		return errors.ErrRedisOperation
	}
	return nil
}

func (r *BaseRedisRepository) Exists(key string) (bool, error) {
	n, err := r.client.Exists(r.ctx, key).Result()
	if err != nil {
		return false, errors.ErrRedisOperation
	}
	return n > 0, nil
}

func (r *BaseRedisRepository) Expire(key string, ttl time.Duration) error {
	if err := r.client.Expire(r.ctx, key, ttl).Err(); err != nil {
		return errors.ErrRedisOperation
	}
	return nil
}
