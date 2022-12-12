package config

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	CTX     = context.Background()
	PodName = getEnv("POD_NAME", "undefined")

	redisPubSubHost     = getEnv("REDIS_PUBSUB_HOST", "localhost")
	redisPubSubPort     = getEnv("REDIS_PUBSUB_PORT", "6379")
	redisPubSubPassword = getEnv("REDIS_PUBSUB_PASSWORD", "")
	PubSubRDB           = redis.NewClient(&redis.Options{
		Addr:     redisPubSubHost + ":" + redisPubSubPort,
		Password: redisPubSubPassword,
		DB:       0,
	})
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
