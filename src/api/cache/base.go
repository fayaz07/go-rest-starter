package cache

import (
	"go-rest-starter/src/core/logger"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func SetRedisClient(c *redis.Client) {
	logger.I("Injecting redis client to cache layer")
	client = c
}
