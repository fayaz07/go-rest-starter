package repository

import "go.mongodb.org/mongo-driver/mongo"

type BaseRepository struct{}

// ---------------------- Database instance
var db *mongo.Database

func UseDb(instance *mongo.Database) {
	db = instance

	initClientColln()
}
