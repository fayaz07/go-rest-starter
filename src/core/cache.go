package core

import (
	"context"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
)

const (
	address   = "127.0.0.1:6379"
	passwd    = ""
	redisDbId = 0
)

var ctx = context.Background()
var _once sync.Once
var redisDb *redis.Client

var redisDbActive bool = false

func InitRedis() {
	_once.Do(func() {
		log.Println("Initializing redis layer...")
		redisDb = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: passwd,    // no password set
			DB:       redisDbId, // use default DB
		})
		if !isRedisHealthy() {
			log.Panic("Redis connection failed!")
		} else {
			log.Println("Connected to Redis!")
		}
	})
}

func isRedisHealthy() bool {
	pingRes := redisDb.Ping(ctx)
	if pingRes != nil && pingRes.String() == "ping: PONG" {
		redisDbActive = true
	} else {
		redisDbActive = false
	}
	return redisDbActive
}
