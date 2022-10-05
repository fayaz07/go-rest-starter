package repository

import (
	"errors"
	"go-rest-starter/src/api/cache"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepository struct {
	model *mongo.Collection
	Name  string
}

// ---------------------- Database instance
var db *mongo.Database

func UseDb(instance *mongo.Database, redisClient *redis.Client) {

	initCache(redisClient)

	db = instance

	initSessionCollection()
	initDeviceColln()
}

func initCache(redisClient *redis.Client) {
	cache.SetRedisClient(redisClient)
}

func id(r *mongo.InsertOneResult) (primitive.ObjectID, error) {
	if oid, ok := r.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	} else {
		return primitive.ObjectID{}, errors.New("invalid mongodb object id")
	}
}
