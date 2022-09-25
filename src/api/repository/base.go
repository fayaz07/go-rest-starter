package repository

import (
	"context"
	"errors"
	"go-rest-starter/src/utils/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

	initSessionCollection()
	initDeviceColln()
}

// Todo: Change this
func GetDbCtx() context.Context {
	return context.Background()
}

func GetRWCtx() *helpers.AppContext {
	return helpers.GetDbRWContext()
}

func id(r *mongo.InsertOneResult) (primitive.ObjectID, error) {
	if oid, ok := r.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	} else {
		return primitive.ObjectID{}, errors.New("invalid mongodb object id")
	}
}
