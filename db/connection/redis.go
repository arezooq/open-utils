package connection

import (
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

func NewRedisClient() *redis.Client {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		// Password: os.Getenv("REDIS_PASSWORD"), // "" if no password
		DB:       db,                          // default 0
		PoolSize: 10,                          // optional: pool size
	})
}
