package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepository struct {
	model *mongo.Collection
	Name  string
}

// ---------------------- Database instance
var db *mongo.Database

func UseDb(instance *mongo.Database) {
	db = instance

	initClientColln()
}

func GetDbCtx() context.Context {
	return context.Background()
}
