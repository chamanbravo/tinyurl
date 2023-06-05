package redis

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func RedisConnection() *redis.Client {
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))
	options := &redis.Options{
		Addr:     os.Getenv("REDIS_SERVER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}

	return redis.NewClient(options)
}
