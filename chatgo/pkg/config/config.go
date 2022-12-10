package config

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	CTX     = context.Background()
	PodName = getEnv("POD_NAME", "undefined")

	redisLeaderHost     = getEnv("REDIS_LEADER_HOST", "localhost")
	redisLeaderPort     = getEnv("REDIS_LEADER_PORT", "6379")
	redisLeaderPassword = getEnv("REDIS_LEADER_PASSWORD", "")
	LeaderRDB           = redis.NewClient(&redis.Options{
		Addr:     redisLeaderHost + ":" + redisLeaderPort,
		Password: redisLeaderPassword,
		DB:       0,
	})

	redisFollowerHost     = getEnv("REDIS_FOLLOWER_HOST", "localhost")
	redisFollowerPort     = getEnv("REDIS_FOLLOWER_PORT", "6379")
	redisFollowerPassword = getEnv("REDIS_FOLLOWER_PASSWORD", "")
	FollowerRDB           = redis.NewClient(&redis.Options{
		Addr:     redisFollowerHost + ":" + redisFollowerPort,
		Password: redisFollowerPassword,
		DB:       0,
	})

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
