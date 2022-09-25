package repository

import (
	"context"
	"go-rest-starter/src/utils/helpers"

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

	initDeviceColln()
}

// Todo: Change this
func GetDbCtx() context.Context {
	return context.Background()
}

func GetRWCtx() *helpers.AppContext {
	return helpers.GetDbRWContext()
}
