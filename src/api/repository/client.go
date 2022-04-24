package repository

import (
	client "go-rest-starter/src/api/models/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type clientRepo BaseRepository

var _clientCollection *mongo.Collection

func initClientColln() {
	_clientCollection = client.GetClientCollection(db)
}

func (clientRepo) CreateClient() {
}
